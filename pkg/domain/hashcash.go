package domain

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"hash"
	"math"
	"strconv"
	"strings"
	"time"
)

// Hashcash provides an implementation of hashcash v1.
type Hashcash struct {
	hasher  hash.Hash // SHA-1
	bits    uint      // Number of zero bits
	zeros   uint      // Number of zero digits
	saltLen uint      // Random salt length
	extra   string    // Extension to add to the minted stamp
}

// New creates a new Hash with specified options.
func NewHashcash(bits uint, saltLen uint, extra string) *Hashcash {
	h := &Hashcash{
		hasher:  sha1.New(),
		bits:    bits,
		saltLen: saltLen,
		extra:   extra}
	h.zeros = uint(math.Ceil(float64(h.bits) / 4.0))
	return h
}

// NewStd creates a new Hash with 20 bits of collision and 8 bytes of salt chars.
func NewStdHashcash() *Hashcash {
	return NewHashcash(20, 8, "")
}

const dateFormat = "060102"

// Mint a new hashcash stamp for resource.
func (h *Hashcash) Mint(resource string) (string, error) {
	salt, err := h.getSalt()
	if err != nil {
		return "", err
	}
	date := time.Now().Format(dateFormat)
	counter := 0
	var stamp string
	for {
		stamp = fmt.Sprintf("1:%d:%s:%s:%s:%s:%x",
			h.bits, date, resource, h.extra, salt, counter)
		if h.checkZeros(stamp) {
			return stamp, nil
		}
		counter++
	}
}

// Check whether a hashcash stamp is valid.
func (h *Hashcash) Check(stamp string) bool {
	if h.checkDate(stamp) {
		return h.checkZeros(stamp)
	}
	return false
}

// CheckNoDate checks whether a hashcash stamp is valid ignoring date.
func (h *Hashcash) CheckNoDate(stamp string) bool {
	return h.checkZeros(stamp)
}

func (h *Hashcash) getSalt() (string, error) {
	buf := make([]byte, h.saltLen)
	n, err := rand.Read(buf)
	_ = n
	if err != nil {
		return "", err
	}
	salt := base64.StdEncoding.EncodeToString(buf)
	return salt[:h.saltLen], nil
}

func (h *Hashcash) checkZeros(stamp string) bool {
	h.hasher.Reset()
	h.hasher.Write([]byte(stamp))
	sum := h.hasher.Sum(nil)
	sumUint64 := binary.BigEndian.Uint64(sum)
	sumBits := strconv.FormatUint(sumUint64, 2)
	zeroes := 64 - len(sumBits)

	return uint(zeroes) >= h.bits
}

func (h *Hashcash) checkDate(stamp string) bool {
	fields := strings.Split(stamp, ":")
	if len(fields) != 7 {
		return false
	}
	then, err := time.Parse(dateFormat, fields[2])
	if err != nil {
		return false
	}
	duration := time.Since(then)
	return duration.Hours()*2 <= 48
}
