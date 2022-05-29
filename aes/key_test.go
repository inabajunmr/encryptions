package aes

import (
	"fmt"
	"testing"
)

func TestSubKeyGen(t *testing.T) {
	// https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197.pdf
	// A.1 Expansion of a 128-bit Cipher Key
	key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
	keys := SubKeyGen(key)
	fmt.Printf("%08x", keys)
}
