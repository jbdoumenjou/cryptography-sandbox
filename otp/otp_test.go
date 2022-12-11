package otp

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/jbdoumenjou/cryptography-sandbox/strings"
)

func TestOTP(t *testing.T) {
	m1 := "attack at dawn"
	c1Hex := "09e1c5f70a65ac519458e7e53f36"

	m1Hex := hex.EncodeToString([]byte(m1))

	key, err := KeyFromPTAndCT(m1Hex, c1Hex)
	if err != nil {
		t.Fail()
	}

	m2 := "attack at dusk"
	m2Hex := hex.EncodeToString([]byte(m2))

	ct2, err := EncryptFromHex(m2Hex, key)
	if err != nil {
		t.Fail()
	}

	fmt.Printf("res: %s\n", ct2)
	ct1, err := EncryptFromHex(m1Hex, key)
	if err != nil {
		t.Fail()
	}

	if ct1 != "09e1c5f70a65ac519458e7e53f36" {
		t.Fail()
	}
	fmt.Printf("m1 xor key:%s\n", ct1)
}

func TestOTP_TwoTimesAttack(t *testing.T) {
	m1 := "attack at dawn"
	m1Hex := hex.EncodeToString([]byte(m1))
	c1Hex := "09e1c5f70a65ac519458e7e53f36"

	m2 := "attack at dusk"
	m2Hex := hex.EncodeToString([]byte(m2))

	r1, err := strings.Xor(m1Hex, m2Hex)
	if err != nil {
		t.Fail()
	}

	r2, err := strings.Xor(r1, c1Hex)
	if err != nil {
		t.Fail()
	}
	if r2 != "09e1c5f70a65ac519458e7d13b31" {
		t.Fail()
	}

	fmt.Printf("%s", r2)
}
