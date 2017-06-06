package cryptostring

import (
	"crypto/rand"
	"encoding/base64"
)

// Cribbed from: https://stackoverflow.com/posts/32351471/revisions

// generateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// SpecifiedLength returns a URL-safe, base64 encoded
// securely generated random string with the specified number of bytes.
func SpecifiedLength(numbytes int) (string, error) {
	b, err := generateRandomBytes(numbytes)
	return base64.URLEncoding.EncodeToString(b), err
}

// String returns a URL-safe, base64 encoded
// securely generated random string with 18 bytes of entropy and 24 characters in length
func String() (string, error) {
	b, err := generateRandomBytes(18)
	return base64.URLEncoding.EncodeToString(b), err
}

// String returns a URL-safe, base64 encoded
// securely generated random string with 18 bytes of entropy and 24 characters in length
// If there's something wrong with the entropy pool, PANIC!
func StringOrBust() string {
	b, err := generateRandomBytes(18)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
