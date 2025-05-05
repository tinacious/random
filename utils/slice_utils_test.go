package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomItemFromList(t *testing.T) {
	result, err := RandomItemFromList([]string{})

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "list cannot be empty", err.Error())
}
