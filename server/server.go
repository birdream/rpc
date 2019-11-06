package server

type RPCServer interface {
	Register(rcvr interface{}, metaData map[string]string) error
	Serve(network string, addr string) error
	Close() error
}
