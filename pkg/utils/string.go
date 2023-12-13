package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenRandomString(length int) (*string, error) {
	// Set the characters to use in the random string
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generate the random string
	b := make([]byte, length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return nil, fmt.Errorf("Failed to generate random number %w", err)
		}
		b[i] = charset[n.Int64()]
	}
	str := string(b)
	return &str, nil
}
