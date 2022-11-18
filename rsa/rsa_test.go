package rsa

import (
	"fmt"
	"testing"
)

func TestRSA(t *testing.T) {
	rsa := NewRSA(2, 5)
	fmt.Printf("%s\n", rsa)

	encrypted := rsa.Encrypt(2)
	if encrypted != 8 {
		t.Fail()
	}

	decrypted := rsa.Decrypt(encrypted)
	if decrypted != 2 {
		t.Fail()
	}
}
