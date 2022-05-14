package des

import "fmt"

// true, true, true, true, false, false, false, false To []byte(0b11110000)
func bitsToBytes(bits []bool) []byte {
	var bytes []byte
	var b byte = 0b00000000
	for i, v := range bits {
		if v {
			b = b | 0b00000001
		}
		if i != 0 {
			if (i+1)%8 == 0 {
				bytes = append(bytes, b)
				b = 0b00000000
			} else if len(bits)-1 == i {
				b = b << (8 - (len(bits))%8)
				bytes = append(bytes, b)
			}
		}
		b = b << 1
	}
	return bytes
}

func printBits(bits []bool) {
	fmt.Print("Bits:")
	for i := 0; i < len(bits); i++ {
		if bits[i] {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		if i != 0 && (i+1)%8 == 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// get bit for index
func bit(index int, bytes []byte) bool {
	switch index % 8 {
	case 0:
		return bytes[index/8]&0b10000000 != 0
	case 1:
		return bytes[index/8]&0b01000000 != 0
	case 2:
		return bytes[index/8]&0b00100000 != 0
	case 3:
		return bytes[index/8]&0b00010000 != 0
	case 4:
		return bytes[index/8]&0b00001000 != 0
	case 5:
		return bytes[index/8]&0b00000100 != 0
	case 6:
		return bytes[index/8]&0b00000010 != 0
	case 7:
		return bytes[index/8]&0b00000001 != 0
	}

	return false
}

// length=6, bytes=0b10101100, result=tftftt
func bytesToBits(length int, bytes []byte) []bool {
	var bits []bool
	for i := 0; i < length; i++ {
		bits = append(bits, bit(i, bytes))
	}
	return bits
}

func shiftLeftRotation(shift int, length int, bytes []byte) []byte {
	bits := bytesToBits(length, bytes)
	shifted := bits[shift:]
	remaining := bits[:shift]
	for i := 0; i < shift; i++ {
		shifted = append(shifted, remaining[i])
	}
	return bitsToBytes(shifted)
}

func xor(a []byte, b []byte) []byte {
	var result []byte
	for i, aval := range a {
		bval := b[i]
		result = append(result, aval^bval)
	}
	return result
}
