package runes

import (
	"strings"
)

// Index returns index number in rune
func Index(s, substr string) int {
	if len(s) < 1 || len(substr) < 1 {
		return -1
	}

	n := strings.Index(s, substr)
	if n < 0 {
		return n
	}
	return Count(s, n)
}

// IndexAll returns the list of all index number in rune
func IndexAll(s, substr string) []int {
	zero := []int{}
	if len(s) < 1 || len(substr) < 1 {
		return zero
	}

	ret := []int{}
	max := len(s)
	l := len(substr)
	pos := 0
	offset := 0
	for {
		n := Index(s[pos:], substr)
		if n < 0 {
			break
		}
		ret = append(ret, n+offset)
		pos += CountByte(s, n) + l // issue: n is rune count, not byte count.
		if pos >= max {
			break
		}
		offset += Count(s, n+l)
	}

	return ret
}
