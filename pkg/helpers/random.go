package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomString() string {
	var result string
	n := 4
	for i := 0; i < n; i++ {
		result = result + randStringRunes(4)
		if i != n-1 {
			result = result + "-"
		}
	}
	return result
}
func initialize() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {

	initialize()

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
