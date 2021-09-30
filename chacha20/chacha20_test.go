package chacha20

import (
	"crypto/rand"
	"testing"
)

func Test(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	nonce := make([]byte, 12)
	rand.Read(nonce)
	v, err := Encrypt([]byte("hellohellohellohellohellohellohellohellohellohellohellohellohell"), key, nonce)
	if err != nil {
		t.Fatal(err)
	}
	v, err = Encrypt(v, key, nonce)
	if err != nil {
		t.Fatal(err)
	}
	if string(v) != "hellohellohellohellohellohellohellohellohellohellohellohellohell" {
		t.Error(string(v))
	}
}
