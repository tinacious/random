package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	length := 27

	result := GeneratePassword(length)

	assert.Equal(t, 27, len(result))
}
