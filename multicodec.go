package mcpacked

import (
	"encoding/binary"
)

type Code uint64

const (
	Unknown = Code(iota)
	Protobuf
	CBOR
	Raw
)

func (c Code) String() string {
	return CodeToString(c)
}

const UnknownMulticodecString = "<Unknown Multicodec>"

func CodeToString(c Code) string {
	switch c {
	case Protobuf:
		return "Protobuf"
	case CBOR:
		return "CBOR"
	case Raw:
		return "Raw"
	default:
		return UnknownMulticodecString
	}
}

func GetCode(data []byte) Code {
	c, _ := binary.Uvarint(data)
	return Code(c)
}

func AddPrefix(c Code, data []byte) []byte {
	buf := make([]byte, len(data)+binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, uint64(c))
	copy(buf[n:], data)
	return buf[:n+len(data)]
}

func SplitPrefix(data []byte) (Code, []byte) {
	c, n := binary.Uvarint(data)
	return Code(c), data[n:]
}
