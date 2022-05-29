package aes

func EncryptBlock(m []byte, k []State) [][]byte {
	n := 10
	// step 2
	s := State{
		m[:4],
		m[4:8],
		m[8:12],
		m[12:16],
	}

	// step 3
	s = addRoundKey(s, k[0])

	for i := 1; /* step 1 */ i <= n; i++ /* step 4e */ {
		// step 4a
		s = SubBytes(s)

		// step 4b
		s = shiftRows(s)

		// step 4c
		if i != n {
			s = mixColumns(s)
		}

		// step 4d
		s = addRoundKey(s, k[i])

	}

	c := s
	return c
}

func mixColumns(s State) State {
	// TODO https://tex2e.github.io/blog/crypto/aes-mix-columns
}

func shiftRows(s State) State {
	r := State{}
	r = append(r, s[0])
	r = append(r, []byte{s[1][1], s[1][2], s[1][3], s[1][0]})
	r = append(r, []byte{s[2][2], s[2][3], s[2][0], s[2][1]})
	r = append(r, []byte{s[3][3], s[3][0], s[3][1], s[3][2]})
	return r
}

func addRoundKey(s, k State) State {
	return State{
		[]byte{s[0][0] ^ k[0][0], s[0][1] ^ k[0][1], s[0][2] ^ k[0][2], s[0][3] ^ k[0][3]},
		[]byte{s[1][0] ^ k[1][0], s[1][1] ^ k[1][1], s[1][2] ^ k[1][2], s[1][3] ^ k[1][3]},
		[]byte{s[2][0] ^ k[2][0], s[2][1] ^ k[2][1], s[2][2] ^ k[2][2], s[2][3] ^ k[2][3]},
		[]byte{s[3][0] ^ k[3][0], s[3][1] ^ k[3][1], s[3][2] ^ k[3][2], s[3][3] ^ k[3][3]},
	}
}
