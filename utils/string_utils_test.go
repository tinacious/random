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
