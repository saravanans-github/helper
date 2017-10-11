package encode

import (
	"encoding/hex"
	"testing"
)

func TestHexStringToBin(t *testing.T) {
	hexConst := "48656c6c6f20476f7068657221" // this is the hex value of "Hello Gopher!"
	s := hex.EncodeToString(HexStringToBin(hexConst))
	if s != hexConst {
		t.FailNow()
	}
}

func TestStringToBase64(t *testing.T) {
	stringConst := "Hello World" // base64 of "Hello World" = SGVsbG8gV29ybGQ=
	if StringToBase64(stringConst) != "SGVsbG8gV29ybGQ=" {
		t.FailNow()
	}
}

// Base64EncodeFromBytes is a base64 encode helper which returns its argument []byte as a base64 encoded string
func TestBytesToBase64(t *testing.T) {
	byteConst := []byte("Hello World") // base64 of "Hello World" = SGVsbG8gV29ybGQ=
	if BytesToBase64(byteConst) != "SGVsbG8gV29ybGQ=" {
		t.FailNow()
	}
}
