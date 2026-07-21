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

// calculateLuminance computes the relative luminance of a color according to WCAG 2.x formulas.
func calculateLuminance(color rgb) float64 {
	// 1. Normalize channels from 0-255 to 0.0-1.0
	r := color.r / 255.0
	g := color.g / 255.0
	b := color.b / 255.0

	// 2. Apply WCAG gamma correction piecewise function
	adjustChannel := func(c float64) float64 {
		if c <= 0.03928 {
			return c / 12.92
		}
		return math.Pow((c+0.055)/1.055, 2.4)
	}

	rAdj := adjustChannel(r)
	gAdj := adjustChannel(g)
	bAdj := adjustChannel(b)

	// 3. Calculate final relative luminance based on human eye sensitivity
	return 0.2126*rAdj + 0.7152*gAdj + 0.0722*bAdj
}

// CalculateContrastRatio calculates the exact WCAG contrast ratio between two hex colors.
func CalculateContrastRatio(hex1, hex2 string) float64 {
	color1 := parseHex(hex1)
	color2 := parseHex(hex2)

	lum1 := calculateLuminance(color1)
	lum2 := calculateLuminance(color2)

	// To calculate contrast, L1 must be the lighter color (higher luminance)
	l1 := math.Max(lum1, lum2)
	l2 := math.Min(lum1, lum2)

	// WCAG Contrast Formula
	return (l1 + 0.05) / (l2 + 0.05)
}

// GetGradeForWCAGAA gives a pass/fail grade for WCAG AA normal text
func GetGradeForWCAGAA(ratio float64) string {
	if ratio >= 4.5 {
		return "PASS"
	} else {
		return "FAIL"
	}
}

// GetGradeForWCAGAALarge gives a pass/fail grade for WCAG AA large text
func GetGradeForWCAGAALarge(ratio float64) string {
	if ratio >= 3.0 {
		return "PASS"
	} else {
		return "FAIL"
	}
}

// GetGradeForWCAGAAA gives a pass/fail grade for WCAG AAA normal text
func GetGradeForWCAGAAA(ratio float64) string {
	if ratio >= 7.0 {
		return "PASS"
	} else {
		return "FAIL"
	}
}

// GetGradeForWCAGAAALarge gives a pass/fail grade for WCAG AAA large text
func GetGradeForWCAGAAALarge(ratio float64) string {
	if ratio >= 4.5 {
		return "PASS"
	} else {
		return "FAIL"
	}
}
