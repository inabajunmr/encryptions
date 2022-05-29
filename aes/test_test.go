package aes

import (
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

func Test2(t *testing.T) {
	fmt.Printf("%02x", 0x38^0xae^0x57)
	t.Error()
}
