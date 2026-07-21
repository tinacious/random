package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStackedResolvedColour_TwoColours(t *testing.T) {
	input := []ColorLayer{
		{
			Hex: "#FFFFFF", Alpha: 1.0,
		},
		{
			Hex: "#111111", Alpha: 0.4,
		},
	}

	result, err := GetStackedResolvedColour(input)
	expected := "#A0A0A0"

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestGetStackedResolvedColour_ThreeColours(t *testing.T) {
	input := []ColorLayer{
		{
			Hex: "#FFFFFF", Alpha: 1.0,
		},
		{
			Hex: "#000000", Alpha: 0.4,
		},
		{
			Hex: "#111111", Alpha: 0.4,
		},
	}

	result, err := GetStackedResolvedColour(input)
	expected := "#636363"

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestGetStackedResolvedColour_TwoColours_Shorthand(t *testing.T) {
	input := []ColorLayer{
		{
			Hex: "#FFF", Alpha: 1.0,
		},
		{
			Hex: "#111", Alpha: 0.4,
		},
	}

	result, err := GetStackedResolvedColour(input)
	expected := "#A0A0A0"

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestGetStackedResolvedColour_ThreeColours_Shorthand(t *testing.T) {
	input := []ColorLayer{
		{
			Hex: "#FFF", Alpha: 1.0,
		},
		{
			Hex: "#000", Alpha: 0.4,
		},
		{
			Hex: "#111", Alpha: 0.4,
		},
	}

	result, err := GetStackedResolvedColour(input)
	expected := "#636363"

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestGetStackedResolvedColour_SemiTransparentWhite(t *testing.T) {
	input := []ColorLayer{
		{
			Hex: "#636363", Alpha: 1.0,
		},
		{
			Hex: "#FFFFFF", Alpha: 0.4,
		},
	}

	result, err := GetStackedResolvedColour(input)
	expected := "#A1A1A1"

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
