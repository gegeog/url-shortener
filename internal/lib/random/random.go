package random

import (
	"math/rand/v2"
	"strings"
	"time"
)

const symbols = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func NewPseudoRandomString(size int) string {
	r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	sb := strings.Builder{}

	for range size {
		sb.WriteByte(symbols[r.IntN(len(symbols))])
	}

	return sb.String()
}
