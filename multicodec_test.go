package mcpacked

import (
	"bytes"
	"testing"
)

func TestEncodeRoundtrip(t *testing.T) {
	data := []byte("Hello World")

	mcdata := AddPrefix(DagProtobuf, data)

	outc, outdata := SplitPrefix(mcdata)
	if outc != DagProtobuf {
		t.Fatal("didnt get same codec as output")
	}

	if GetCode(mcdata) != DagProtobuf {
		t.Fatal("GetCode returned incorrect code")
	}

	if !bytes.Equal(outdata, data) {
		t.Fatal("output data not the same as input data")
	}
}

func TestStringer(t *testing.T) {
	if DagCBOR.String() != "dag-cbor" {
		t.Fatal("stringify failed")
	}

	if DagProtobuf.String() != "dag-pb" {
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
