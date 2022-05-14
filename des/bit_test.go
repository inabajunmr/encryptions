package des

import (
	"testing"
)

func TestBit(t *testing.T) {
	bytes := []byte{0b10101010, 0b01010101}
	if bit(0, bytes) != true {
		t.Errorf("%08b", bytes)
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

func TestBitsToBytes1(t *testing.T) {
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

func TestBitsToBytes2(t *testing.T) {
	result := bitsToBytes([]bool{true, false, true, false})
	if len(result) != 1 {
		t.Error(result)
	}

	if result[0] != 0b10100000 {
		t.Errorf("%b", result[0])
	}
}

func TestBitsToBytes3(t *testing.T) {
	result := bitsToBytes([]bool{true, false, true, false, true})
	if len(result) != 1 {
		t.Error(result)
	}

	if result[0] != 0b10101000 {
		t.Errorf("%b", result[0])
	}
}

func TestShiftLeft(t *testing.T) {
	result := shiftLeftRotation(3, 15, []byte{0b10101111, 0b01010010})
	if len(result) != 2 {
		t.Error(result)
	}
	if result[0] != 0b01111010 {
		t.Errorf("%08b %08b", result[0], result[1])
	}
	if result[1] != 0b10011010 {
		t.Errorf("%08b %08b", result[0], result[1])
	}
}

func TestBytesToBits(t *testing.T) {
	result := bytesToBits(12, []byte{0b10101010, 0b01010000})
	if len(result) != 12 {
		t.Error(result)
	}
	if result[0] != true {
		t.Error(result)
	}
	if result[1] != false {
		t.Error(result)
	}
	if result[2] != true {
		t.Error(result)
	}
	if result[3] != false {
		t.Error(result)
	}
	if result[4] != true {
		t.Error(result)
	}
	if result[5] != false {
		t.Error(result)
	}
	if result[6] != true {
		t.Error(result)
	}
	if result[7] != false {
		t.Error(result)
	}
	if result[8] != false {
		t.Error(result)
	}
	if result[9] != true {
		t.Error(result)
	}
	if result[10] != false {
		t.Error(result)
	}
	if result[11] != true {
		t.Error(result)
	}
}

func TestXor(t *testing.T) {
	result := xor([]byte{0b10101010, 0b01010101}, []byte{0b11110000, 0b00001111})
	if len(result) != 2 {
		t.Error(result)
	}
	if result[0] != 0b01011010 {
		t.Errorf("%b", result[0])
	}
	if result[1] != 0b01011010 {
		t.Errorf("%b", result[1])
	}
}
