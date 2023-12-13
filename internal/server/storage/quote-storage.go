package storage

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
)

//go:embed quotes.txt
var rawQuotes []byte

type QuoteStorage struct {
	quotes [][]byte
}

func NewQuoteStorage() *QuoteStorage {
	return &QuoteStorage{quotes: bytes.Split(rawQuotes, []byte("\n\n"))}
}

func (s QuoteStorage) GetRandom() (*string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(s.quotes))))
	if err != nil {
		return nil, fmt.Errorf("failed to generate random number %w", err)
	}
	q := string(s.quotes[n.Int64()])
	return &q, nil
}
