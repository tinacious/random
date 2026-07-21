package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Colour generation logic from: https://stackoverflow.com/a/46909816/1870884

const letterBytes = "abcdef0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func RandomColour() string {
	return RandStringBytesMaskImprSrc(6)
}

func RandomColours(n int) []string {
	var colours = make([]string, n)
	for i := 0; i < n; i++ {
		colours[i] = RandomColour()
	}
	return colours
}

// ColorLayer represents a single layer in the compositing stack.
type ColorLayer struct {
	Hex   string
	Alpha float64 // 0.0 to 1.0
}

// rgb represents internal color channels for calculation
type rgb struct {
	r, g, b float64
}

// parseHex converts a hex string (e.g., "#FFF" or "#FFFFFF") to an rgb struct.
func parseHex(hex string) rgb {
	hex = strings.TrimPrefix(hex, "#")

	// Handle 3-character shorthand (e.g., "FFF" -> "FFFFFF")
	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}

	// Parse the hex string into a 32-bit integer
	parsed, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return rgb{0, 0, 0} // Fallback for invalid hex
	}

	return rgb{
		r: float64((parsed >> 16) & 255),
		g: float64((parsed >> 8) & 255),
		b: float64(parsed & 255),
	}
}

// GetStackedResolvedColour calculates the final visible color from a stack of layers.
func GetStackedResolvedColour(layers []ColorLayer) (string, error) {
	if len(layers) == 0 {
		return "", fmt.Errorf("no colours in the layer list")
	}

	// The bottom layer acts as our starting solid background
	base := parseHex(layers[0].Hex)

	// Blend each subsequent layer on top iteratively
	for _, layer := range layers[1:] {
		top := parseHex(layer.Hex)
		alpha := layer.Alpha

		base.r = (top.r * alpha) + (base.r * (1.0 - alpha))
		base.g = (top.g * alpha) + (base.g * (1.0 - alpha))
		base.b = (top.b * alpha) + (base.b * (1.0 - alpha))
	}

	// Round to the nearest integer and format back to a 6-character uppercase Hex string
	resolvedColour := fmt.Sprintf("#%02X%02X%02X",
		int(math.Round(base.r)),
		int(math.Round(base.g)),
		int(math.Round(base.b)),
	)

	return resolvedColour, nil
}
