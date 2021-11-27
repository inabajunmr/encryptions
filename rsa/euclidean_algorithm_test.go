package rsa

import "testing"

func TestGcd(t *testing.T) {
	r := gcd(1, 6)
	if r != 1 {
		t.Error()
	}

	r = gcd(2, 6)
	if r != 2 {
		t.Error()
	}

	r = gcd(36, 24)
	if r != 12 {
		t.Error()
	}
}
