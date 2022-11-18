package rsa

import (
	"fmt"
	"math"
)

// RSA contains data needed to apply RSA algorithm.
type RSA struct {
	// first prime number
	p int
	// second prime number
	q int
	// prime numbers product
	n int
	// Euler Totient
	t int
	// Encode number
	e int
	// Decode number
	d int
}

// NewRSA creates a RSA struct used to Encrypt and Decrypt.
// This is a very naive approach and works only for very small number
// TODO: use a math lib with better approach
func NewRSA(p, q int) *RSA {
	r := &RSA{p: p, q: q}
	r.n = p * q
	r.t = (p - 1) * (q - 1)

	for i := 2; i < r.t; i++ {
		if coprime(r.t, i) && coprime(r.n, i) {
			r.e = i
			break
		}
	}

	var results []int
	for i := 1; len(results) <= 2; i++ {
		if (i*r.e)%r.t == 1 {
			results = append(results, i)
		}
	}
	// arbitrary take the second result
	r.d = results[1]

	return r
}

// String implements Stringer.
func (r *RSA) String() string {
	res := fmt.Sprintf("p:%d, q:%d\n", r.p, r.q)
	res += fmt.Sprintf("n = %d = p x q = %d x %d\n", r.n, r.p, r.q)
	res += fmt.Sprintf("t = %d = (p-1)(q-1) = (%d-1)(%d-1)\n", r.t, r.p, r.q)
	res += fmt.Sprintf("e = %d (coprime with t=%d and n=%d)\n", r.e, r.t, r.n)
	res += fmt.Sprintf("d = %d, (e.d mod t = 1)\n", r.d)
	res += fmt.Sprintf("public key (e=%d,n=%d)\n", r.e, r.n)
	res += fmt.Sprintf("private key (d=%d,n=%d)\n", r.d, r.n)

	return res
}

// Encrypt encrypts an integer following RSA algorithm.
// This is a very naive approach and works only for very small number
// TODO: use a math lib with better approach
func (r *RSA) Encrypt(m int) int {
	pow := math.Pow(float64(m), float64(r.e))
	fmt.Printf("%d^%d mod %d\n", m, r.e, r.n)

	return int(pow) % r.n
}

// Decrypt decrypts an integer following RSA algorithm.
// This is a very naive approach and works only for very small number
// TODO: use a math lib with better approach
func (r *RSA) Decrypt(m int) int {
	pow := math.Pow(float64(m), float64(r.d))
	fmt.Printf("%d^%d mod %d\n", m, r.d, r.n)

	return int(pow) % r.n
}

// coprime checks if two integers are co-primes.
func coprime(a, b int) bool {
	// Set a to gcd(a,b)
	var t int
	for b != 0 {
		t = b
		b = a % b
		a = t
	}

	// By definition, a and b are co-prime if gcd(a,b) == 1
	return a == 1
}
