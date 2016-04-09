package sunrpc

import "testing"
import "encoding/hex"

func Assert(t *testing.T, tag string, got []byte, exp string) {
	if hex.EncodeToString(got) != exp {
		t.Error(tag, exp, got)
	}
}

func TestMarshal(t *testing.T) {
	var enc uint32 = 255
	var xdr = make([]byte, 4)
	Marshal(enc, xdr)
	Assert(t, "uint32", xdr, "000000ff")
}

func TestUnmarshal(t *testing.T) {

}
