package encode

import (
	"encoding/hex"
	"testing"
)

func TestGetHexStringAsBin(t *testing.T) {
	hexConst := "48656c6c6f20476f7068657221" // this is the hex value of "Hello Gopher!"
	s := hex.EncodeToString(GetHexStringAsBin(hexConst))
	if s != hexConst {
		t.FailNow()
	}
}

func TestBase64Encode(t *testing.T) {
	stringConst := "Hello World" // base64 of "Hello World" = SGVsbG8gV29ybGQ=
	if Base64Encode(stringConst) != "SGVsbG8gV29ybGQ=" {
		t.FailNow()
	}
}

// Base64EncodeFromBytes is a base64 encode helper which returns its argument []byte as a base64 encoded string
func TestBase64EncodeFromBytes(t *testing.T) {
	byteConst := []byte("Hello World") // base64 of "Hello World" = SGVsbG8gV29ybGQ=
	if Base64EncodeFromBytes(byteConst) != "SGVsbG8gV29ybGQ=" {
		t.FailNow()
	}
}
