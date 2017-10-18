package crypto

import (
	"helper/encode"
	"reflect"
	"testing"
)

func TestPadWithPKCS5(t *testing.T) {
	padConst := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	paddedConst := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 6, 6, 6, 6, 6}
	if !reflect.DeepEqual(PadWithPKCS5(padConst), paddedConst) {
		t.FailNow()
	}
}

func TestCreateSha1Hash(t *testing.T) {
	shaConst := []byte("Hello World")
	constAfterShaing := encode.HexStringToBin("0a4d55a8d778e5022fab701977c5d840bbc486d0")
	if !reflect.DeepEqual(CreateSha1Hash(shaConst), constAfterShaing) {
		t.FailNow()
	}
}

func TestEncryptWithAesCBC(t *testing.T) {
	clearConst := []byte("1234567890123456")
	encryptConst := encode.HexStringToBin("1bb4f4175f942fbcaf4211e7f866c310")
	keyConstInBin := encode.HexStringToBin("dfbc480b4333ea1895c27470ae5dab79")
	ivConstInBin := encode.HexStringToBin("efd140d77b5eb2050d610da5dcba22f4")

	encryptedData, err := EncryptWithAesCBC(keyConstInBin, ivConstInBin, clearConst)
	if !reflect.DeepEqual(encryptedData, encryptConst) || err != nil {
		t.FailNow()
	}
}

func TestEncryptWithAesCBC_Negative_NoPadding(t *testing.T) {
	clearConst := []byte("123456789")
	keyConstInBin := encode.HexStringToBin("dfbc480b4333ea1895c27470ae5dab79")
	ivConstInBin := encode.HexStringToBin("efd140d77b5eb2050d610da5dcba22f4")

	_, err := EncryptWithAesCBC(keyConstInBin, ivConstInBin, clearConst)
	if err.Error() != _CryptoErrorDataNotPadded {
		t.FailNow()
	}
}

func TestEncryptWithAesCBC_Negative_InvalidKeyIv(t *testing.T) {
	clearConst := []byte("1234567890123456")
	keyConstInBin := encode.HexStringToBin("dfbc480b4333ea1895c27470ae5dab79")
	ivConstInBin := encode.HexStringToBin("efd140d77b5eb2050d610da5dcba22f4")

	inValidKeyConstInBin := encode.HexStringToBin("dfbc480b4333ea1895c27470ae5dab")
	invalidIvConstInBin := encode.HexStringToBin("efd140d77b5eb2050d610da5dcba22")

	var err error

	_, err = EncryptWithAesCBC(inValidKeyConstInBin, ivConstInBin, clearConst)
	if err == nil {
		t.FailNow()
	}

	_, err = EncryptWithAesCBC(keyConstInBin, invalidIvConstInBin, clearConst)
	if err.Error() != _CryptoErrorInvalidIvLength {
		t.FailNow()
	}
}
