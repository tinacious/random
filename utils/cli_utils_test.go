package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntRangeFromString_withValidRange(t *testing.T) {
	input := "1-3"
	start, end, err := IntRangeFromString(input, 1, 5)

	assert.Nil(t, err)
	assert.Equal(t, 1, start)
	assert.Equal(t, 3, end)
}

func TestIntRangeFromString_withInvalidRange(t *testing.T) {
	input := "hello-world"
	start, end, err := IntRangeFromString(input, 1, 100)

	assert.NotNil(t, err)
	assert.Equal(t, "strconv.Atoi: parsing \"hello\": invalid syntax", err.Error())
	assert.Equal(t, -1, start)
	assert.Equal(t, -1, end)
}

func TestIntRangeFromString_whenStartRangeIsTooLow_returnsMinValue(t *testing.T) {
	input := "1023-1024"
	start, end, err := IntRangeFromString(input, 1024, 1024)

	assert.Nil(t, err)
	assert.Equal(t, 1024, start)
	assert.Equal(t, 1024, end)
}

func TestIntRangeFromString_whenStartRangeIsTooHigh_returnsMaxValue(t *testing.T) {
	input := "66535-66666"
	start, end, err := IntRangeFromString(input, 66535, 66535)

	assert.Nil(t, err)
	assert.Equal(t, 66535, start)
	assert.Equal(t, 66535, end)
}
