package mcpacked

import (
	"bytes"
	"testing"
)

func TestEncodeRoundtrip(t *testing.T) {
	data := []byte("Hello World")

	mcdata := AddPrefix(Protobuf, data)

	outc, outdata := SplitPrefix(mcdata)
	if outc != Protobuf {
		t.Fatal("didnt get same codec as output")
	}

	if GetCode(mcdata) != Protobuf {
		t.Fatal("GetCode returned incorrect code")
	}

	if !bytes.Equal(outdata, data) {
		t.Fatal("output data not the same as input data")
	}
}

func TestStringer(t *testing.T) {
	if CBOR.String() != "CBOR" {
		t.Fatal("stringify failed")
	}

	if Protobuf.String() != "Protobuf" {
		t.Fatal("stringify failed")
	}

	if Code(125125).String() != UnknownMulticodecString {
		t.Fatal("expected unknown mcodec string for random value")
	}
}

func TestEdgeCases(t *testing.T) {
	c := GetCode(nil)
	if c != Unknown {
		t.Fatal("invalid buffer should return Unknown")
	}
}
