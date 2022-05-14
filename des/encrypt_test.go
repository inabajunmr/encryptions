package des

import "testing"

func TestEncrypt(t *testing.T) {
	subKeys := SubKeyGen([]byte("1234abcd"))
	t.Log(Encrypt([]byte("superman"), subKeys))
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
