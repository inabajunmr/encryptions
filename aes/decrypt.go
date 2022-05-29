package aes

func DecryptBlock(m []byte, k []State) State {
	n := 10
	// step 2
	s := FromByte(m)

	// step 3
	for i := n; /* step 1 */ i >= 1; i-- /* step 3e */ {
		// step 3a
		s = addRoundKey(s, k[i])

		// step 3b
		if i != n {
			s = invMixColumns(s)
		}

		// step 3c
		s = invShiftRows(s)

		// step 3d
		s = InvSubBytes(s)

	}

	// step 4
	s = addRoundKey(s, k[0])

	c := s
	return c
}

// https://tex2e.github.io/blog/crypto/aes-mix-columns
func invMixColumns(s State) State {
	for i := 0; i < 4; i++ {
		var t [4]byte
		t[0] = dot(0x0e, s[0][i]) ^ dot(0x0b, s[1][i]) ^ dot(0x0d, s[2][i]) ^ dot(0x09, s[3][i])
		t[1] = dot(0x09, s[0][i]) ^ dot(0x0e, s[1][i]) ^ dot(0x0b, s[2][i]) ^ dot(0x0d, s[3][i])
		t[2] = dot(0x0d, s[0][i]) ^ dot(0x09, s[1][i]) ^ dot(0x0e, s[2][i]) ^ dot(0x0b, s[3][i])
		t[3] = dot(0x0b, s[0][i]) ^ dot(0x0d, s[1][i]) ^ dot(0x09, s[2][i]) ^ dot(0x0e, s[3][i])
		s[0][i] = t[0]
		s[1][i] = t[1]
		s[2][i] = t[2]
		s[3][i] = t[3]
	}
	return s
}

func invShiftRows(s State) State {
	r := State{}
	r = append(r, s[0])
	r = append(r, []byte{s[1][3], s[1][0], s[1][1], s[1][2]})
	r = append(r, []byte{s[2][2], s[2][3], s[2][0], s[2][1]})
	r = append(r, []byte{s[3][1], s[3][2], s[3][3], s[3][0]})
	return r
}
