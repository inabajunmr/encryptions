package des

import (
	"errors"
)

func Encrypt(plain []byte, key []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, errors.New("key size must be 8 bytes")
	}
	return encryptDecrypt(plain, SubKeyGen(key)), nil
}

func Decrypt(encrypted []byte, key []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, errors.New("key size must be 8 bytes")
	}
	subKeys := SubKeyGen(key)
	var reverseSubKeys [][]byte
	for i := 15; i >= 0; i-- {
		reverseSubKeys = append(reverseSubKeys, subKeys[i])
	}
	return encryptDecrypt(encrypted, reverseSubKeys), nil
}

func encryptDecrypt(target []byte, keys [][]byte) []byte {
	var result []byte
	for i := 0; i < len(target); i += 8 {
		var block []byte
		for j := 0; j < 8; j++ {
			if len(target) > i+j {
				block = append(block, target[i+j])
			} else {
				block = append(block, 0x00)
			}
		}
		result = append(result, EncryptBlock(block, keys)...)
	}
	return result
}
