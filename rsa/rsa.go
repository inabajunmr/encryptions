package rsa

func encrypt(plain string, e int, n int) string {
	bytes := []byte(plain)
	var encrypted []byte
	for _, b := range bytes {
		encrypted = append(encrypted, byte(int(b)^e%n))
	}
	return string(encrypted)

}

func decrypt(encrypted string, key rsa_key) string {
	l := lcm(key.p-1, key.q-1)
	d := 0
	y := 0
	for {
		temp := d*key.e - y*l
		if temp == 1 {
			break
		}
		if temp < 1 {
			d++
		} else {
			y++
		}
	}

	var plain []byte
	for _, v := range []byte(encrypted) {
		plain = append(plain, byte(int(v)^d%key.n))
	}

	return string(plain)

}

func lcm(val1 int, val2 int) int {
	gcd := Gcd(val1, val2)
	return int(float64(val1) / float64(val2) * float64(gcd))
}

type rsa_key struct {
	p, q, e, n int
}

func gen_keys(prime1 int, prime2 int) *rsa_key {
	p := prime1
	q := prime2
	n := p * q
	e := 2
	return &rsa_key{p, q, n, int(e)}
}
