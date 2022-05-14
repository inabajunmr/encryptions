package des

import (
	"reflect"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("password")
	plain := []byte("helloworld123456")
	enc, err := Encrypt(plain, key)
	if err != nil {
		t.Error(err)
	}
	dec, err := Decrypt(enc, key)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(plain, dec) {
		t.Errorf("plain and dec must be same. dec:%v\n", string(dec))
	}
}
