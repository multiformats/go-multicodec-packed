package mcpacked

import (
	"encoding/binary"
)

type Code uint64

const (
	Unknown     = Code(0)
	Git         = Code(0x69)
	DagProtobuf = Code(0x70)
	DagCBOR     = Code(0x71)
	Raw         = Code(0x55)

	EthereumBlock = Code(0x90)
	EthereumTx    = Code(0x91)
	BitcoinBlock  = Code(0xb0)
	BitcoinTx     = Code(0xb1)
	ZcashBlock    = Code(0xc0)
	ZcashTx       = Code(0xc1)
)

func (c Code) String() string {
	return CodeToString(c)
}

const UnknownMulticodecString = "<Unknown Multicodec>"

func CodeToString(c Code) string {
	switch c {
	case Git:
		return "git"
	case DagProtobuf:
		return "dag-pb"
	case DagCBOR:
		return "dag-cbor"
	case Raw:
		return "bin"
	case BitcoinBlock:
		return "bitcoin-block"
	case BitcoinTx:
		return "bitcoin-tx"
	case EthereumBlock:
		return "eth-block"
	case EthereumTx:
		return "eth-tx"
	case ZcashBlock:
		return "zcash-block"
	case ZcashTx:
		return "zcash-tx"
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
