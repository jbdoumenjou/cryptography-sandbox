package otp

import (
	"fmt"

	"github.com/jbdoumenjou/cryptography-sandbox/strings"
)

// EncryptFromHex encrypts the given message with the given key
// such as ct = m xor k
// The message and the key must have the same len
func EncryptFromHex(m, k string) (string, error) {
	s, err := strings.Xor(m, k)
	if err != nil {
		return "", fmt.Errorf("xor: %w", err)
	}

	return s, nil
}

// KeyFromPTAndCT deduces the key from the plaintext message and the ciphertext
// knowing ct = m xor k => k = m xor ct
// The plaintext and cipher text must be from the same type (hexa/ascii)...
func KeyFromPTAndCT(pt, ct string) (string, error) {
	return EncryptFromHex(pt, ct)
}
