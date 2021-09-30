package chacha20

import (
	"encoding/binary"
	"errors"
	"unsafe"
)

func Encrypt(plain []byte, key []byte, nonce []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("unexpected key length")
	}
	if len(nonce) != 12 {
		return nil, errors.New("unexpected nonce length")
	}

	var counter uint32
	encrypted := []byte{}
	for i := 0; i < len(plain)/64; i++ {
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, counter)
		stream := keyStream(key, nonce, bytes)
		if len(plain) <= i*64 {
			break
		}
		for j := 0; j < 64; j++ {
			if len(plain) > i*64+j {
				s := stream[j]
				p := plain[i*64+j]
				encrypted = append(encrypted, p^s)
			}
		}
	}

	return encrypted, nil
}

func keyStream(key []byte, nonce []byte, counter []byte) []byte {

	state := []uint32{0x61707865, 0x3320646e, 0x79622d32, 0x6b206574}

	state = append(state, binary.BigEndian.Uint32(key[0:4]))
	state = append(state, binary.BigEndian.Uint32(key[4:8]))
	state = append(state, binary.BigEndian.Uint32(key[8:12]))
	state = append(state, binary.BigEndian.Uint32(key[12:16]))

	state = append(state, binary.BigEndian.Uint32(key[16:20]))
	state = append(state, binary.BigEndian.Uint32(key[20:24]))
	state = append(state, binary.BigEndian.Uint32(key[24:28]))
	state = append(state, binary.BigEndian.Uint32(key[28:32]))

	state = append(state, binary.BigEndian.Uint32(counter))
	state = append(state, binary.BigEndian.Uint32(nonce[0:4]))
	state = append(state, binary.BigEndian.Uint32(nonce[4:8]))
	state = append(state, binary.BigEndian.Uint32(nonce[8:12]))

	for i := 0; i < 10; i++ {
		round(state)
	}

	stream := []byte{}
	for _, s := range state {
		stream = append(stream, (*[4]byte)(unsafe.Pointer(&s))[:]...)
	}

	return stream
}

func round(state []uint32) []uint32 {
	state[0], state[4], state[8], state[12] = qr(state[0], state[4], state[8], state[12])
	state[1], state[5], state[9], state[13] = qr(state[1], state[5], state[9], state[13])
	state[2], state[6], state[10], state[14] = qr(state[2], state[6], state[10], state[14])
	state[3], state[7], state[11], state[15] = qr(state[3], state[7], state[11], state[15])
	state[0], state[5], state[10], state[15] = qr(state[0], state[5], state[10], state[15])
	state[1], state[6], state[11], state[12] = qr(state[1], state[6], state[11], state[12])
	state[2], state[7], state[8], state[13] = qr(state[2], state[7], state[8], state[13])
	state[3], state[4], state[9], state[14] = qr(state[3], state[4], state[9], state[14])
	return state
}

func qr(a uint32, b uint32, c uint32, d uint32) (uint32, uint32, uint32, uint32) {
	a = a + b
	d = d ^ a
	d = rotLeft(d, 16)

	c = c + d
	b = b ^ c
	b = rotLeft(b, 12)

	a = a + b
	d = d ^ a
	d = rotLeft(d, 8)

	c = c + d
	b = b ^ c
	b = rotLeft(b, 7)
	return a, b, c, d
}

func rotLeft(v uint32, o int) uint32 {
	return v>>o | v<<(32-o)
}
