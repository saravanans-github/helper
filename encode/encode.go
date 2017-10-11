package encode

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

// HexStringToBin is a string helper which interprets its argument string as a hex number and returns it as a []byte
func HexStringToBin(s string) []byte {
	h, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return h
}

// StringToBase64 base64 encode helper which returns its argument string as a base64 encoded string
func StringToBase64(s string) string {
	return BytesToBase64([]byte(s))
}

// BytesToBase64 is a base64 encode helper which returns its argument []byte as a base64 encoded string
func BytesToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
