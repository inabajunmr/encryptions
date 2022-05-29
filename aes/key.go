package aes

// 128 bit key
func SubKeyGen(key []byte) [][][]byte {
	// step 1
	w := [][]byte{key[0:4], key[4:8], key[8:12], key[12:16]}

	// step 3
	k := [][][]byte{{w[0], w[1], w[2], w[3]}}

	// step 4
	for i := 1; /* step 2 */ i <= 10; i++ /* step 4e */ {
		n := 4*i - 1

		// step 4a-1
		rotWord := rotWord(w[n])

		// step 4a-2
		var subWord []byte
		for wi := range rotWord {
			subWord = append(subWord, SubByte(rotWord[wi]))
		}

		// step 4a-3
		var temp []byte
		for ri := range rcon[i-1] {
			temp = append(temp, rcon[i-1][ri]^subWord[ri])
		}

		// step 4c
		for j := 0; /* step 4b */ j < 4; j++ /* step 4c-3 */ {
			// step 4c-1
			var xor []byte
			for tempi := range temp {
				xor = append(xor, w[4*(i-1)+j][tempi]^temp[tempi])
			}
			w = append(w, xor)

			// step 4c-2
			temp = w[4*i+j]
		}
		k = append(k, [][]byte{w[4*i], w[4*i+1], w[4*i+2], w[4*i+3]})
	}

	return k
}

func rotWord(w []byte) []byte {
	return []byte{w[1], w[2], w[3], w[0]}
}

var rcon = [][]byte{
	{0x01, 0, 0, 0},
	{0x02, 0, 0, 0},
	{0x04, 0, 0, 0},
	{0x08, 0, 0, 0},
	{0x10, 0, 0, 0},
	{0x20, 0, 0, 0},
	{0x40, 0, 0, 0},
	{0x80, 0, 0, 0},
	{0x1b, 0, 0, 0},
	{0x36, 0, 0, 0},
}
