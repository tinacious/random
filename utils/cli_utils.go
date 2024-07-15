package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func IntRangeFromString(input string, min int, max int) (int, int, error) {
	splits := strings.Split(input, "-")

	if len(splits) != 2 {
		return -1, -1, fmt.Errorf("invalid format: %s", input)
	}

	start, err := strconv.Atoi(splits[0])
	if err != nil {
		return -1, -1, err
	}

	end, err := strconv.Atoi(splits[1])
	if err != nil {
		return -1, -1, err
	}

	if end <= start {
		return -1, -1, fmt.Errorf("end must be larger than start")
	}

	start = int(math.Max(float64(min), float64(start)))
	end = int(math.Min(float64(max), float64(end)))

	return start, end, nil
}
