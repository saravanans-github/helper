package crypto

import (
	"bytes"
	"crypto/aes"
)

// PadWithPKCS5 is a crypto helper which pads its argument byte per the PKCS#5 standard
func PadWithPKCS5(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}
