package server

import (
	"context"
	"errors"
	"io"
	"log"
	"reflect"
	"rpc/codec"
	"rpc/protocol"
	"rpc/transport"
	"strings"
	"sync"
)

type simpleServer struct {
	codec      codec.Codec
	serviceMap sync.Map
	tr         transport.ServerTransport
	mutex      sync.Mutex
	shutdown   bool

	option Option
}

type service struct {
	name    string
	typ     reflect.Type
	rcvr    reflect.Value
	methods map[string]*methodType
}

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
}

func NewSimpleServer(option Option) RPCServer {
	s := new(simpleServer)
	s.option = option
	s.codec = codec.GetCodec(option.SerializeType)

	return s
}

func (s *simpleServer) writeErrorResponse(response *protocol.Message, w io.Writer, err string) {
	response.Error = err
	log.Println(response.Error)
	response.StatusCode = protocol.StatusError
	response.Data = response.Data[:0]
	_, _ = w.Write(protocol.EncodeMessage(s.option.ProtocolType, response))
}

func (s *simpleServer) serveTransport(tr transport.Transport) {
	for {
		request, err := protocol.DecodeMessage(s.option.ProtocolType, tr)

		if err != nil {
			if err == io.EOF {
				log.Printf("client has closed this connection: %s", tr.RemoteAddr().String())
			} else if strings.Contains(err.Error(), "use of closed network connection") {
				log.Printf("rpcx: connection %s is closed", tr.RemoteAddr().String())
			} else {
				log.Printf("rpcx: failed to read request: %v", err)
			}

			return
		}

		response := request.Clone()
		response.MessageType = protocol.MessageTypeResponse

		sname := request.ServiceName
		mname := request.MethodName
		srvInterface, ok := s.serviceMap.Load(sname)

		if !ok {
			s.writeErrorResponse(response, tr, "can not find service")
			return
		}

		srv, ok := srvInterface.(*service)
		if !ok {
			s.writeErrorResponse(response, tr, "not *service type")
			return

		}

		mtype, ok := srv.methods[mname]
		if !ok {
			s.writeErrorResponse(response, tr, "can not find method")
			return
		}
		argv := newValue(mtype.ArgType)
		replyv := newValue(mtype.ReplyType)

		ctx := context.Background()
		err = s.codec.Decode(request.Data, argv)

		var returns []reflect.Value
		if mtype.ArgType.Kind() != reflect.Ptr {
			returns = mtype.method.Func.Call([]reflect.Value{srv.rcvr,
				reflect.ValueOf(ctx),
				reflect.ValueOf(argv).Elem(),
				reflect.ValueOf(replyv)})
		} else {
			returns = mtype.method.Func.Call([]reflect.Value{srv.rcvr,
				reflect.ValueOf(ctx),
				reflect.ValueOf(argv),
				reflect.ValueOf(replyv)})
		}
		if len(returns) > 0 && returns[0].Interface() != nil {
			err = returns[0].Interface().(error)
			s.writeErrorResponse(response, tr, err.Error())
			return
		}

		responseData, err := codec.GetCodec(request.SerializeType).Encode(replyv)
		if err != nil {
			s.writeErrorResponse(response, tr, err.Error())
			return
		}

		response.StatusCode = protocol.StatusOK
		response.Data = responseData

		_, err = tr.Write(protocol.EncodeMessage(s.option.ProtocolType, response))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (s *simpleServer) Register(rcvr interface{}, metaData map[string]string) error {
	typ := reflect.TypeOf(rcvr)
	name := typ.Name()

	srv := new(service)
	srv.name = name
	srv.rcvr = reflect.ValueOf(rcvr)
	srv.typ = typ
	methods := suitableMethods(typ, true)
	srv.methods = methods

	if len(srv.methods) == 0 {
		var errorStr string

		method := suitableMethods(reflect.PtrTo(srv.typ), false)
		if len(method) != 0 {
			errorStr = "rpcx.Register: type " + name + " has no exported methods of suitable type (hint: pass a pointer to value of that type)"

		} else {
			errorStr = "rpcx.Register: type " + name + " has no exported methods of suitable type"
		}

		log.Println(errorStr)
		return errors.New(errorStr)
	}

	if _, duplicate := s.serviceMap.LoadOrStore(name, srv); duplicate {
		return errors.New("rpc: service already defined: " + name)
	}

	return nil
}

func (s *simpleServer) Serve(network string, addr string) (err error) {
	s.tr = transport.NewServerTransport(s.option.TransportType)
	if err = s.tr.Listen(network, addr); err != nil {
		log.Println(err)
		return
	}

	var conn transport.Transport

	for {
		if conn, err = s.tr.Accept(); err != nil {
			log.Println(err)
			return
		}

		go s.serveTransport(conn)
	}

	return nil
}

func (s *simpleServer) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.shutdown = true

	err := s.tr.Close()

	s.serviceMap.Range(func(key, value interface{}) bool {
		s.serviceMap.Delete(key)
		return true
	})
	return err
}
