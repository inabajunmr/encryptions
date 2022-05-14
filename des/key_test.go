package des

import (
	"testing"
)

func Test(t *testing.T) {
	// key := []byte{0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111}
	c, _ := NewCipher([]byte("1234abcd"))
	d := c.(*desCipher).subkeys
	for i, v := range d {
		t.Logf("%v:%08b\n", i, v)
	}
}

func TestSubKeygen(t *testing.T) {
	subKeys := SubKeyGen([]byte("1234abcd"))
	for i, v := range subKeys {
		t.Logf("%v: %08b\n", i, v)
	}
}

func TestPg1(t *testing.T) {
	c0, c1 := pg1([]byte("abcd1234"))
	if len(c0) != 4 {
		t.Errorf("%08b %08b %08b %08b", c0[0], c0[1], c0[2], c0[3])
	}
	if len(c1) != 4 {
		t.Errorf("%08b %08b %08b %08b", c1[0], c1[1], c1[2], c1[3])
	}
}
