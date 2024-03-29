package aes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncryptBlock(t *testing.T) {
	input := []byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34}
	fmt.Printf("plain: %02x \n", input)

	key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}

	encrypted := EncryptBlock(input, SubKeyGen(key))
	encrypted.Print()
	fmt.Printf("enc: %02x \n", encrypted.ToByte())

	decrypted := DecryptBlock(encrypted.ToByte(), SubKeyGen(key))
	fmt.Printf("dec: %02x \n", decrypted.ToByte())

	if !reflect.DeepEqual(input, decrypted.ToByte()) {
		t.Error("plain and decrypted are not same.")
	}
}
