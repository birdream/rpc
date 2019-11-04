package protocol

// MessageType ..
type MessageType byte

// CompressType ..
type CompressType byte

// StatusCode ..
type StatusCode byte

// ProtocolType ..
type ProtocolType byte

//请求类型
const (
	MessageTypeRequest MessageType = iota
	MessageTypeResponse

	CompressTypeNone CompressType = iota

	StatusOK StatusCode = iota
	StatusError

	Default ProtocolType = iota

	RequestSeqKey     = "rpc_request_seq"
	RequestTimeoutKey = "rpc_request_timeout"
	MetaDataKey       = "rpc_meta_data"
)
