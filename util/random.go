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

func RandomInt(min int32, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomFloat() float64 {
	return rand.Float64()
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

func RandomMoney() float64 {
	return RandomFloat() * 100
}

func RandomCurreny() string {
	curreny := []string{"EUR", "USD"}
	l := len(curreny)
	return curreny[rand.Intn(l)]
}
