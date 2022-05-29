package aes

import (
	"encoding/binary"
)

func EncryptBlock(m []byte, k []State) State {
	n := 10
	// step 2
	s := FromByte(m)

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

// https://tex2e.github.io/blog/crypto/aes-mix-columns
func mixColumns(s State) State {
	for i := 0; i < 4; i++ {
		var t [4]byte
		t[0] = dot(2, s[0][i]) ^ dot(3, s[1][i]) ^ s[2][i] ^ s[3][i]
		t[1] = s[0][i] ^ dot(2, s[1][i]) ^ dot(3, s[2][i]) ^ s[3][i]
		t[2] = s[0][i] ^ s[1][i] ^ dot(2, s[2][i]) ^ dot(3, s[3][i])
		t[3] = dot(3, s[0][i]) ^ s[1][i] ^ s[2][i] ^ dot(2, s[3][i])
		s[0][i] = t[0]
		s[1][i] = t[1]
		s[2][i] = t[2]
		s[3][i] = t[3]
	}
	return s
}

// https://tex2e.github.io/blog/crypto/aes-mix-columns
func dot(x, y byte) byte {
	product := byte(0)
	for mask := byte(0x01); mask != 0x00; mask <<= 1 {
		if y&mask != 0x00 {
			product ^= x
		}
		x = xtime(x)
	}

	return product
}

// https://tex2e.github.io/blog/crypto/aes-mix-columns
func xtime(b byte) byte {
	data := binary.LittleEndian.Uint16([]byte{b, 0})
	x := data << 1
	r := make([]byte, 8)
	if x&0x100 != 0 {
		binary.LittleEndian.PutUint16(r, x^0x1b)
	} else {
		binary.LittleEndian.PutUint16(r, x)
	}
	return r[0]
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
