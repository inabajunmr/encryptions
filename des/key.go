package des

func SubKeyGen(key []byte) [][]byte {
	// step 1
	c0, d0 := pg1(key)

	// step 2 & 3
	c := [][]byte{c0}
	d := [][]byte{d0}
	var keys [][]byte
	for i := 1; i <= 16; i++ /* step 3c */ {
		// step 3a
		cbefore := c[i-1]
		dbefore := d[i-1]
		var shiftedC []byte
		var shiftedD []byte
		if i == 1 || i == 2 || i == 9 || i == 16 {
			// shift 1
			shiftedC = shiftLeftRotation(1, 28, cbefore)
			shiftedD = shiftLeftRotation(1, 28, dbefore)
		} else {
			// shift 2
			shiftedC = shiftLeftRotation(2, 28, cbefore)
			shiftedD = shiftLeftRotation(2, 28, dbefore)
		}
		c = append(c, shiftedC)
		d = append(d, shiftedD)

		// step 3b
		keys = append(keys, pg2(c[i], d[i]))
	}

	// step 4
	return keys
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
		bits = append(bits, bit(PG1_ARRAY[i]-1, key))
	}
	printBits(bits)
	return bitsToBytes(bits[:28]), bitsToBytes(bits[28:])
}

var PG2_ARRAY = []int{
	14, 17, 11, 24, 1, 5,
	3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8,
	16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55,
	30, 40, 51, 45, 33, 48,
	44, 49, 39, 56, 34, 53,
	46, 42, 50, 36, 29, 32}

func pg2(c []byte, d []byte) []byte {
	merged := append(c, d...)
	var bits []bool
	for i := 0; i < 48; i++ {
		bits = append(bits, bit(PG2_ARRAY[i]-1, merged))
	}
	return bitsToBytes(bits)
}
