package aes

import "fmt"

// 4 * 4(128 bit)
type State [][]byte

func (s *State) Print() {
	for _, v := range *s {
		for _, vv := range v {
			fmt.Printf("%02x ", vv)
		}
		fmt.Println()
	}
}

func (s *State) ToByte() []byte {
	return []byte{
		(*s)[0][0], (*s)[1][0], (*s)[2][0], (*s)[3][0],
		(*s)[0][1], (*s)[1][1], (*s)[2][1], (*s)[3][1],
		(*s)[0][2], (*s)[1][2], (*s)[2][2], (*s)[3][2],
		(*s)[0][3], (*s)[1][3], (*s)[2][3], (*s)[3][3],
	}
}

func FromByte(m []byte) State {
	return State{
		{m[0], m[4], m[8], m[12]},
		{m[1], m[5], m[9], m[13]},
		{m[2], m[6], m[10], m[14]},
		{m[3], m[7], m[11], m[15]},
	}
}

func FromByte2(v [][]byte) State {
	return State{
		{v[0][0], v[1][0], v[2][0], v[3][0]},
		{v[0][1], v[1][1], v[2][1], v[3][1]},
		{v[0][2], v[1][2], v[2][2], v[3][2]},
		{v[0][3], v[1][3], v[2][3], v[3][3]},
	}
}
