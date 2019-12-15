package runes

import (
	"strings"
)

// Index returns index number in rune
func Index(s, substr string) int {
	n := strings.Index(s, substr)
	if n < 0 {
		return n
	}
	return Count(s, n)
}

func IndexAll(s, substr string) []int {
	return nil
}
