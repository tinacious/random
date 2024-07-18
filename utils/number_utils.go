package utils

import (
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RandomNumberBetweenRange(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(max-min+1) + min
}

func FormatNumberWithDelimiter(num int, delimiter string) string {
	p := message.NewPrinter(language.English)

	formattedNumber := p.Sprintf("%d", num)
	if delimiter != "," {
		formattedNumber = strings.ReplaceAll(formattedNumber, ",", delimiter)
	}

	return formattedNumber
}
