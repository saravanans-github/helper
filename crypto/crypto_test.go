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

func TestEncryptWithAesCbc(t *testing.T) {
	clearConst := []byte("1234567890123456")
	encryptConst := encode.HexStringToBin("1bb4f4175f942fbcaf4211e7f866c310")
	keyConstInBin := encode.HexStringToBin("dfbc480b4333ea1895c27470ae5dab79")
	ivConstInBin := encode.HexStringToBin("efd140d77b5eb2050d610da5dcba22f4")

	encryptedData, err := EncryptWithAesCbc(keyConstInBin, ivConstInBin, clearConst)
	if !reflect.DeepEqual(encryptedData, encryptConst) || err != nil {
		t.FailNow()
	}
}

func TestEncryptWithAesCbc_Negative_NoPadding(t *testing.T) {
	clearConst := []byte("123456789")
	keyConstInBin := encode.HexStringToBin("dfbc480b4333ea1895c27470ae5dab79")
	ivConstInBin := encode.HexStringToBin("efd140d77b5eb2050d610da5dcba22f4")

	_, err := EncryptWithAesCbc(keyConstInBin, ivConstInBin, clearConst)
	if err.Error() != _CryptoErrorDataNotPadded {
		t.FailNow()
	}
}

func TestEncryptWithAesCbc_Negative_InvalidKeyIv(t *testing.T) {
	clearConst := []byte("1234567890123456")
	keyConstInBin := encode.HexStringToBin("dfbc480b4333ea1895c27470ae5dab79")
	ivConstInBin := encode.HexStringToBin("efd140d77b5eb2050d610da5dcba22f4")

	inValidKeyConstInBin := encode.HexStringToBin("dfbc480b4333ea1895c27470ae5dab")
	invalidIvConstInBin := encode.HexStringToBin("efd140d77b5eb2050d610da5dcba22")

	var err error

	_, err = EncryptWithAesCbc(inValidKeyConstInBin, ivConstInBin, clearConst)
	if err == nil {
		t.FailNow()
	}

	_, err = EncryptWithAesCbc(keyConstInBin, invalidIvConstInBin, clearConst)
	if err.Error() != _CryptoErrorInvalidIvLength {
		t.FailNow()
	}
}

func TestDecryptWithRC4_64(t *testing.T) {
	keyConstInBin := encode.HexStringToBin("c6592459cd8cb8e7c8fd6c36bb0f0092ce3779e14d61338fe5a36c945170f85d")
	dataConstInBin := encode.HexStringToBin("88583ef721601bbfc58d01169e8821c31ea28c1d58d2d81920d2be7d751dbd0a7642233b20a4dba163b8e1d4d2daf467fd355d1c57abde011a909fa8a87a01c2")
	decryptedDataConst := "1ae8ccd0e7985cc0b6203a55855a1034afc252980e970ca90e5202689f947ab9"

	decryptedData, err := DecryptWithRC4_64(keyConstInBin, dataConstInBin)
	if err != nil {
		t.FailNow()
	}

	if string(decryptedData) != decryptedDataConst {
		t.FailNow()
	}
}
