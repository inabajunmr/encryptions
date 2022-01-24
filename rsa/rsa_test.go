package rsa

import (
	"testing"
)

func TestRsa(t *testing.T) {
	key := gen_keys(3, 5)
	enc := encrypt("hello world", key.e, key.n)
	dec := decrypt(enc, *key)
	if dec != "hello world" {
		t.Error(dec)
	}
}
