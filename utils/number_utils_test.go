package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomNumberBetweenRange(t *testing.T) {
	for x := 1; x <= 100; x++ {
		result := RandomNumberBetweenRange(10, 15)

		assert.LessOrEqual(t, 10, result)
		assert.GreaterOrEqual(t, 15, result)
	}
}

func TestFormatNumberWithDelimiter(t *testing.T) {
	assert.Equal(t, "123,456,789", FormatNumberWithDelimiter(123456789, ","))
	assert.Equal(t, "123.456.789", FormatNumberWithDelimiter(123456789, "."))
	assert.Equal(t, "9,999,999", FormatNumberWithDelimiter(9999999, ","))
	assert.Equal(t, "9 999 999", FormatNumberWithDelimiter(9999999, " "))
}
