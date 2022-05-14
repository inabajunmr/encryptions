package des

import (
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	// genkey
	key := []byte("1234abcd")
	subKeys := SubKeyGen(key)
	t.Logf("subKeys: %v", subKeys)

	// plain
	plain := []byte("superman")
	t.Logf("plain  :%v", plain)

	// encrypt
	encrypt := Encrypt(plain, subKeys)
	t.Logf("encrypt:%08b", encrypt)

	// gen decrypt key
	var reverseSubKeys [][]byte
	for i := 15; i >= 0; i-- {
		reverseSubKeys = append(reverseSubKeys, subKeys[i])
	}
	t.Logf("reverseSubKeys: %v", reverseSubKeys)

	// decrypt
	decrypt := Encrypt([]byte("superman"), reverseSubKeys)
	t.Logf("decrypt:%v", decrypt)

	// assertion
	if !reflect.DeepEqual(plain, decrypt) {
		t.Error("plain and decrypt aren't same.")
	}
}

func TestSBox(t *testing.T) {
	result1 := sbox(bytesToBits(6, []byte{0b01100100}), S_BOX1)
	if result1[0] != true || result1[1] != false || result1[2] != false || result1[3] != true {
		t.Error(result1)
	}
	result2 := sbox(bytesToBits(6, []byte{0b01100100}), S_BOX5)
	if result2[0] != false || result2[1] != false || result2[2] != true || result2[3] != true {
		t.Error(result1)
	}
	result3 := sbox(bytesToBits(6, []byte{0b10110000}), S_BOX6)
	if result3[0] != true || result3[1] != true || result3[2] != false || result3[3] != false {
		t.Error(result1)
	}
}
