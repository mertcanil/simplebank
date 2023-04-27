package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max)
}

func RandomString(n int) string {
	var s strings.Builder
	l := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(l)]
		s.WriteByte(c)
	}
	return s.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 100)
}

func RandomCurreny() string {
	curreny := []string{"EUR", "USD"}
	l := len(curreny)
	return curreny[rand.Intn(l)]
}
