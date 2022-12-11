package strings

import (
	"errors"
)

// Xor computes the binary xor of n strings of the same length.
func Xor(s1, s2 string) (string, error) {
	l := len(s1)
	if l != len(s2) {
		return "", errors.New("message and key must have the same len")
	}

	res := make([]uint8, l)
	for i := 0; i < l; i++ {
		res[i] = s1[i] ^ s2[i]
	}

	return string(res), nil
}
