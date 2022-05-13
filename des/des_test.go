package des

import (
	"testing"
)

func TestSubKeygen(t *testing.T) {
	// 64 bit secret key
	// key := []byte("abcd1234")
	// subKeys := SubKeyGen(key)

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

func TestBit(t *testing.T) {
	bytes := []byte{0b10101010, 0b01010101}
	if bit(0, bytes) != true {
		t.Error()
	}
	if bit(1, bytes) != false {
		t.Error()
	}
	if bit(2, bytes) != true {
		t.Error()
	}
	if bit(3, bytes) != false {
		t.Error()
	}
	if bit(4, bytes) != true {
		t.Error()
	}
	if bit(5, bytes) != false {
		t.Error()
	}
	if bit(6, bytes) != true {
		t.Error()
	}
	if bit(7, bytes) != false {
		t.Error()
	}
	if bit(8, bytes) != false {
		t.Error()
	}
	if bit(9, bytes) != true {
		t.Error()
	}
	if bit(10, bytes) != false {
		t.Error()
	}
	if bit(11, bytes) != true {
		t.Error()
	}
	if bit(12, bytes) != false {
		t.Error()
	}
	if bit(13, bytes) != true {
		t.Error()
	}
	if bit(14, bytes) != false {
		t.Error()
	}
	if bit(15, bytes) != true {
		t.Error()
	}
}

func TestBitsToBytes(t *testing.T) {
	result := bitsToBytes([]bool{true, false, true, false, true, true, false, true,
		true, false, false, false, true, true, true, false})
	if len(result) != 2 {
		t.Error(result)
	}

	if result[0] != 0b10101101 {
		t.Errorf("%b", result[0])
	}
	if result[1] != 0b10001110 {
		t.Errorf("%b\n", result[1])
	}
}
