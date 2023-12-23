package pow

import (
	"fmt"
	"testing"
	"time"
)

// TODO : complete test cases
func TestNewPoW(t *testing.T) {
	tm := time.Now()
	p := NewPoW([]byte("secret"), []byte("nonce"), 6, tm)
	b, _ := p.Generate()

	sp := NewPoW(nil, nil, 0, time.Now())
	sp.SetDifficulty(2)
	sp.SetNonce("nonce")
	sp.SetSecret("secret")
	sp.SetTimestamp(tm)

	if s, ok := sp.Solve(b); ok {
		if s != p.GetA() {
			t.Fail()
		}
		fmt.Println(s, b)
	}

}
