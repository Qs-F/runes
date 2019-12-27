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

// IndexAll returns the list of all index number in rune
func IndexAll(s, substr string) []int {
	ret := []int{}
	l := Count(substr, -1)

	pos := 0
	for {
		n := Index(s[pos:], substr)
		if n < 0 {
			break
		}
		ret = append(ret, n+pos)
		pos += n + l
	}

	return ret
}
