package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type (
	PoW interface {
	}
	pow struct {
		A          hashBytes
		B          hashBytes
		Secret     []byte
		Difficulty int
		Nonce      []byte
		TimeStamp  TimeStamp
	}
	TimeStamp []byte
	hashBytes []byte
)

var arr = strings.Split("0123456789abcdefghijklmnopqrstuvwxyz", "")

// NewPoW initializes a new instance of the PoW struct.
//
// It takes the following parameters:
//   - secret: a byte slice containing the secret value.
//   - nonce: a byte slice containing the nonce value.
//   - difficulty: an integer representing the difficulty level.
//   - timestamp: a time.Time object representing the timestamp.
//
// It returns a PoW interface.
func NewPoW(secret []byte, nonce []byte, difficulty int, timestamp time.Time) PoW {
	tb := make([]byte, 8)
	binary.LittleEndian.PutUint64(tb, uint64(timestamp.Unix()))
	return &pow{
		nil,
		nil,
		secret,
		difficulty,
		nonce,
		tb,
	}
}

// SetDifficulty updates the difficulty in pow struct
func (p *pow) SetDifficulty(level int) {
	if level >= 64 {
		panic("Dude! it's impossible")
	}
	p.Difficulty = level
}

// SetNonce updates the nonce in pow struct
func (p *pow) SetNonce(str string) {
	p.Nonce = []byte(str)
}

// SetSecret updates the secret in pow struct
func (p *pow) SetSecret(secret string) {
	p.Secret = []byte(secret)
}

// SetTimestamp updates the timestamp in pow struct
func (p *pow) SetTimestamp(t time.Time) {
	bts := make([]byte, 8)
	binary.LittleEndian.PutUint64(bts, uint64(t.Unix()))
	p.TimeStamp = bts
}

// Generate generates the problem to solve
func (p *pow) Generate() (string, string) {
	p.A = p.hash(p.Secret, p.TimeStamp, p.Nonce)
	p.B = p.hash(p.A, p.TimeStamp, p.Nonce)
	a := p.A.String()
	return p.B.String(), a[:len(a)-(p.Difficulty)]
}

// Solve solves the given problem
func (p *pow) Solve(str string) (string, bool) {
	arrLen := 64 - len(str)
	return p.backtrack(arrLen, str, "")
}

func (p *pow) backtrack(arrLen int, str string, path string) (string, bool) {
	if len(path) == arrLen {
		fmt.Sprintf("%s%s\n", str, path)
		if ok, _ := p.isOK(str, path); ok {
			return fmt.Sprintf("%s%s", str, path), true
		}
		return path, false
	}
	for _, value := range arr {
		if s, ok := p.backtrack(arrLen, str, fmt.Sprintf("%s%s", path, value)); ok {
			return fmt.Sprintf("%s", s), true
		}
	}
	return "", false
}
func (p *pow) hash(data ...[]byte) []byte {
	h := sha256.New()
	for i := 0; i < len(data); i++ {
		h.Write(data[i])
	}
	return h.Sum(nil)
}

func (p *pow) isOK(a string, g string) (bool, error) {
	ds, err := hex.DecodeString(fmt.Sprintf("%s%s", a, g))
	if err != nil {
		return false, err
	}
	return bytes.Equal(p.hash(ds, p.TimeStamp, p.Nonce), p.B), nil
}

func (h hashBytes) String() string {
	return hex.EncodeToString(h)
}
