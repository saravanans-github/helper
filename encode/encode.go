package encode

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

// GetHexStringAsBin is a string helper which interprets its argument string as a hex number and returns it as a []byte
func GetHexStringAsBin(s string) []byte {
	h, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return h
}

// Base64Encode base64 encode helper which returns its argument string as a base64 encoded string
func Base64Encode(s string) string {
	return Base64EncodeFromBytes([]byte(s))
}

// Base64EncodeFromBytes is a base64 encode helper which returns its argument []byte as a base64 encoded string
func Base64EncodeFromBytes(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
