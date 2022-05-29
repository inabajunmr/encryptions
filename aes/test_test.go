package aes

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	x2 := xtime(0x57)
	fmt.Printf("%02x\n", x2)
	x3 := xtime(x2)
	fmt.Printf("%02x\n", x3)
	x4 := xtime(x3)
	fmt.Printf("%02x\n", x4)
	x5 := xtime(x4)
	fmt.Printf("%02x\n", x5)
	x6 := xtime(x5)
	fmt.Printf("%02x\n", x6)
	x7 := xtime(x6)
	fmt.Printf("%02x\n", x7)
	x8 := xtime(x7)
	fmt.Printf("%02x\n", x8)
	t.Error()
}

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

func Test2(t *testing.T) {
	fmt.Printf("%02x", 0x38^0xae^0x57)
	t.Error()
}
