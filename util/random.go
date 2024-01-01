package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt returns a random integer between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandomString returns a random string of length n
func RandomString(n int) string {
	sb := strings.Builder{}

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[RandomInt(0, len(alphabet))])
	}

	return sb.String()
}
