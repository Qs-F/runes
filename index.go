package runes

import (
	"strings"
	"unicode/utf8"
)

func Index(s, substr string) int {
	n := strings.Index(s, substr)
	if n < 0 {
		return n
	}
	return utf8.RuneCountInString(s[:n])
}

func Next(s string, bytepos int) int {
	// 6 is the max value of utf8 byte
	for count := 1; count < 7; count++ {
		if utf8.RuneStart(s[bytepos+count]) {
			return bytepos + count
		}
	}
	return -1
}
