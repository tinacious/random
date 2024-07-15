package utils

import (
	"math/rand"
	"time"
)

func RandomNumberBetweenRange(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(max-min+1) + min
}
