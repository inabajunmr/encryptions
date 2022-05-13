package des

func SubKeyGen(key []byte) [][]byte {
	// step 1
	c0, d0 := pg1(key)

	// step 2 & 3
	c := [][]byte{c0}
	d := [][]byte{d0}
	for i := 1; i <= 16; i++ /* step 3c */ {
		// step 3a
		cbefore := c[i-1]
		dbefore := d[i-1]
		var shiftedC []byte
		var shiftedD []byte
		if i == 1 || i == 2 || i == 9 || i == 16 {
			// shift 1
			shiftedC = cbefore // TODO shift
			shiftedD = dbefore

		} else {
			// shift 2
			shiftedC = cbefore // TODO shift
			shiftedD = dbefore
		}
		c[i] = shiftedC
		d[i] = shiftedD

		// step 3b
	}

	// step 4
	return nil
}

var PG1_ARRAY = []int{
	57, 49, 41, 33, 25, 17, 9,
	1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27,
	19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15,
	7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 28, 20, 12, 4}

func pg1(key []byte) ([]byte /* C0 */, []byte /* C1 */) {
	var bits []bool
	for i := 0; i < 56; i++ {
		bits = append(bits, bit(PG1_ARRAY[i], key))
	}
	return bitsToBytes(bits[:28]), bitsToBytes(bits[29:])
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
				bytes = append(bytes, b)
			}
		}
		b = b << 1
	}
	return bytes
}
