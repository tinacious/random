package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// RandomItemFromList
func RandomItemFromList[T any](items []T) (*T, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("list cannot be empty")
	}

	// Initialize the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Random index
	idx := rng.Intn(len(items))

	return &items[idx], nil
}

// ProperNounCase changes the case to use proper noun case in English, e.g. JOHN -> John, sally -> Sally
func ProperNounCase(input string) string {
	if len(input) == 0 {
		return input
	}

	r := []rune(strings.ToLower(input))
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}
