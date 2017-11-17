package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rc4"
	"crypto/sha1"
	"fmt"
)

const _CryptoErrorDataNotPadded = "The data to be encrypted is not padded."
const _CryptoErrorCreatingKey = "There was an error creating the key. [%s]"
const _CryptoErrorInvalidIvLength = "There was an error with the IV length. It should be 16 bytes"
const _CryptoErrorDecryptingWithRC4 = "There was an error decrypting. [%s]"

type Crypto struct {
	errorMessage string
}

// PadWithPKCS5 is a crypto helper which pads its argument byte per the PKCS#5 standard
func PadWithPKCS5(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// CreateSha1Hash is a crypto helper which creates a SHA1 hash of the argument passed in
func CreateSha1Hash(data []byte) []byte {
	// step 1: generate a sha1 hash of the data
	h := sha1.New()
	h.Write(data)
	dataAfterShaing := h.Sum(nil)

	return dataAfterShaing
}

// EncryptWithAesCbc is a crypto helper which encrypts the argument using AES CBC cipher.
func EncryptWithAesCbc(key []byte, iv []byte, data []byte) ([]byte, error) {

	// Ensure the data is padded
	if len(data)%aes.BlockSize != 0 {
		return []byte{}, &Crypto{errorMessage: _CryptoErrorDataNotPadded}
	}

	// Length of IV is not checked by the library... so we check it
	lengthOfIv := len(iv)
	if lengthOfIv != aes.BlockSize {
		return []byte{}, &Crypto{errorMessage: _CryptoErrorInvalidIvLength}
	}

	// Create the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, &Crypto{errorMessage: fmt.Sprintf(_CryptoErrorCreatingKey, err.Error())}
	}

	// Encrypt the data using now
	mode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(data))
	mode.CryptBlocks(cipherText, data)

	return cipherText, nil
}

// DecryptWithRC4_64 is a crypto helper which decrypts the argument RC4-64 bit cipher.
func DecryptWithRC4_64(key []byte, data []byte) ([]byte, error) {
	// to decode, you will need to initialize a new cipher with the same key.
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return []byte{}, &Crypto{errorMessage: fmt.Sprintf(_CryptoErrorDecryptingWithRC4, err.Error())}
	}

	cipher.XORKeyStream(data, data)
	return data, err
}

func (c *Crypto) Error() string {
	return c.errorMessage
}
