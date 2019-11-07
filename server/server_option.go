package server

import (
	"rpc/codec"
	"rpc/protocol"
	"rpc/transport"
)

type Option struct {
	ProtocolType  protocol.ProtocolType
	SerializeType codec.SerializeType
	CompressType  protocol.CompressType
	TransportType transport.TransportType
}

var DefaultOption = Option{
	ProtocolType:  protocol.Default,
	SerializeType: codec.MessagePack,
	CompressType:  protocol.CompressTypeNone,
	TransportType: transport.TCPTransport,
}

// defaultOption = defaultOption
