package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var currencies = [...]string{
	"USD",
	"EUR",
	"JPY",
	"SGD",
	"HKD"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of english alphabets of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomAccountName generates an random account name
func RandomAccountName() string {
	return RandomString(5)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
