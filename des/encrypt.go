package des

import "fmt"

func Encrypt(m []byte, keys [][]byte) []byte {
	// step 2
	ip := ip(m)

	// step 3
	r := [][]byte{ip[:4]}
	l := [][]byte{ip[4:]}

	// step 4
	for n := 1; /* step 1 */ 1 <= 16; n++ /* step 4d */ {
		// step 4a
		y := roundFunction(r[n-1], keys[n-1])
		// step 4b
		ln := xor(l[n-1], y)
		rn := r[n-1]
		// step 4c
		l = append(l, rn)
		r = append(r, ln)
	}

	// step 5
	c := ip2(append(l[15], r[15]...))
	return c

}

func ip2(b []byte) []byte {
	panic("unimplemented")
}

func roundFunction(x, key []byte) []byte {
	// step 1
	xd := e(x)

	// step 2
	xdd := xor(xd, key)

	// step 3
	xddd := s(xdd)

	// step 4
	y := p(xddd)

	// step 5
	return y
}

func p(xddd []byte) []byte {
	// TODO
	panic("unimplemented")
}

var S_BOX1 = [][]byte{
	{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
	{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
	{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
	{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
}

var S_BOX2 = [][]byte{
	{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
	{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
	{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
	{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
}

var S_BOX3 = [][]byte{
	{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
	{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
	{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
	{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
}

var S_BOX4 = [][]byte{
	{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
	{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
	{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
	{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
}

var S_BOX5 = [][]byte{
	{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
	{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
	{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
	{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
}

var S_BOX6 = [][]byte{
	{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
	{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
	{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
	{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
}

var S_BOX7 = [][]byte{
	{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
	{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
	{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
	{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
}

var S_BOX8 = [][]byte{
	{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
	{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
	{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
	{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
}

var S_BOXES = [][][]byte{
	S_BOX1, S_BOX2, S_BOX3, S_BOX4, S_BOX5, S_BOX6, S_BOX7, S_BOX8,
}

func s(xdd []byte) []byte {
	// split each 6 bits
	bits := bytesToBits(48, xdd)
	var bitss [][]bool
	for i := 0; i <= 7; i++ {
		bitss = append(bitss, bits[i*6:(i+1)*6])
	}

	var xddd []bool
	for i, b := range bitss {
		xddd = append(xddd, sbox(b, S_BOXES[i])...)
	}

	return bitsToBytes(xddd)
}

func sbox(b []bool, sbox [][]byte) []bool {
	row := 0
	if b[5] {
		row += 1
	}
	if b[0] {
		row += 2
	}

	col := 0
	if b[4] {
		col += 1
	}
	if b[3] {
		col += 2
	}
	if b[2] {
		col += 4
	}
	if b[1] {
		col += 8
	}
	fmt.Printf("row: %v col: %v b: %v sboxval: %v", row, col, b, sbox[row][col])
	return bytesToBits(8, []byte{sbox[row][col]})[4:]
}

var E_ARRAY = []int{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1,
}

func e(x []byte) []byte {
	var bits []bool
	for i := 0; i < 48; i++ {
		bits = append(bits, bit(E_ARRAY[i], x))
	}
	return bitsToBytes(bits)
}

var IP1_ARRAY = []int{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

func ip(key []byte) []byte {
	var bits []bool
	for i := 0; i < 64; i++ {
		bits = append(bits, bit(IP1_ARRAY[i], key))
	}
	return bitsToBytes(bits)
}
