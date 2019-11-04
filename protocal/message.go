package protocol

import "io"

//-----------------------------------------------------------------------------------------------------
//|2byte |1byte   |4byte        |4byte         | header length |(total length - header length - 4byte)|
//-----------------------------------------------------------------------------------------------------
//|magic |version |total length |header length |     header    |                    body              |
//-----------------------------------------------------------------------------------------------------

// Protocol 定义了如何构造和序列化一个完整的消息体
type Protocol interface {
	NewMessage() *Message
	DecodeMessage(r io.Reader) (*Message, error)
	EncodeMessage(message *Message) []byte
}

var protocols = map[ProtocolType]Protocol{
	Default: &RPCProtocol{},
}

// Header ..
type Header struct {
	Seq           uint64              // 序号, 用来唯一标识请求或响应
	MessageType   MessageType         // 消息类型，用来标识一个消息是请求还是响应
	CompressType  CompressType        // 压缩类型，用来标识一个消息的压缩方式
	SerializeType codec.SerializeType // 序列化类型，用来标识消息体采用的编码方式
	StatusCode    StatusCode          // 状态类型，用来标识一个请求是正常还是异常
	ServiceName   string              // 服务名
	MethodName    string              // 方法名
	Error         string              // 方法调用发生的异常
	MetaData      map[string]string   // 其他元数据
}

// Message 表示一个消息体
type Message struct {
	*Header        //head部分
	Data    []byte //body部分
}

func NewMessage(t ProtocolType) *Message {
	return protocols[t].NewMessage()
}

func DecodeMessage(t ProtocolType, r io.Reader) (*Message, error) {
	return protocols[t].DecodeMessage(r)
}

func EncodeMessage(t ProtocolType, m *Message) []byte {
	return protocols[t].EncodeMessage(m)
}

type Message struct {
	*Header
	Data []byte
}

func (m Message) Clone() *Message {
	header := *m.Header
	c := new(Message)
	c.Header = &header
	c.Data = m.Data
	return c
}
