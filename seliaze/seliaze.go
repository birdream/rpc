package seliaze

import "github.com/vmihailenco/msgpack"

// Codec ..
type Codec interface {
	Encode(value interface{}) ([]byte, error)
	Decode(data []byte, value interface{}) error
}

type SerializeType byte

const (
	MessagePack SerializeType = iota
)

var codecs = map[SerializeType]Codec{
	MessagePack: &MessagePackCodec{},
}

// NewCodec ..
func NewCodec(t SerializeType) Codec {
	return codecs[t]
}

// MessagePackCodec ..
type MessagePackCodec struct{}

// Encode ..
func (c MessagePackCodec) Encode(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// Decode ..
func (c MessagePackCodec) Decode(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
