package helman

import (
	"math/rand"
)

func Exchange() {
	var alicePrivateKey = rand.Int()
	var bobPrivateKey = rand.Int()

	var p = 29
	var g = 2

	var alicePublicKey = g ^ alicePrivateKey%p
	var bobPublicKey = g ^ bobPrivateKey%p

	var sharedKeyForAlice = alicePublicKey ^ bobPrivateKey%p
	var sharedKeyForBob = bobPublicKey ^ alicePrivateKey%p

	if sharedKeyForAlice != sharedKeyForBob {
		panic("something wrong")
	}
}
