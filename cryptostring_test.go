package cryptostring

import (
	"testing"
)

func TestSpecifiedLength(t *testing.T) {
	str, err := SpecifiedLength(16)
	if err != nil {
		t.Error(err)
	}
	if len(str) != 24 {
		t.Error("Length does not match")
	}

	str, err = SpecifiedLength(20)
	if err != nil {
		t.Error(err)
	}
	if len(str) != 28 {
		t.Error("Length does not match")
	}
}

func TestString(t *testing.T) {
	str, err := String()
	if err != nil {
		t.Error(err)
	}
	if len(str) != 24 {
		t.Error("Length does not match")
	}
}

const n = 18

func BenchmarkCryptoString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SpecifiedLength(n)
	}
}
