package utils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomHexWithLength(t *testing.T) {
	result, err := RandomHex(16)

	fmt.Println(result)
	assert.Nil(t, err)
	assert.Equal(t, 32, len(result))
}

func TestTrimEmptyLinesRemovesEmptyLines(t *testing.T) {
	input := `one

two

three
`
	assert.Equal(t, 6, len(strings.Split(input, "\n")))

	result := TrimEmptyLines(input)

	assert.Equal(t, 3, len(strings.Split(result, "\n")))
}

func TestProperNounCase_FromUpper(t *testing.T) {
	input := "BOB"
	result := ProperNounCase(input)

	assert.Equal(t, "Bob", result)
}

func TestProperNounCase_FromLower(t *testing.T) {
	input := "john"
	result := ProperNounCase(input)

	assert.Equal(t, "John", result)
}

func TestProperNounCase_WithNoChanges(t *testing.T) {
	input := "Devin"
	result := ProperNounCase(input)

	assert.Equal(t, "Devin", result)
}

func TestProperNounCase_Empty(t *testing.T) {
	input := ""
	result := ProperNounCase(input)

	assert.Equal(t, "", result)
}
