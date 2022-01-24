package rsa

func Gcd(a int, b int) int {

	l := a
	r := b

	for {
		nl := r
		r = l % r
		if r == 0 {
			return nl
		}
		l = nl
	}
}
