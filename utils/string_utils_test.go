package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomHexWithLength(t *testing.T) {
	result, err := RandomHex(16)

	fmt.Println(result)
	assert.Nil(t, err)
	assert.Equal(t, 32, len(result))
}
