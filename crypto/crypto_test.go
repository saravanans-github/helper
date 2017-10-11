package crypto

import (
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
