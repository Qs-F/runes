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

func Next(s string, bytepos int) int {
	// 6 is the max value of utf8 byte
	for count := 0; count < 7; count++ {
		if utf8.RuneStart(s[bytepos+count+1]) {
			return bytepos + count + 1
		}
	}
	return -1
}

func Before(s string, bytepos int) int {
	// 6 is the max value of utf8 byte
	for count := 0; count < 7; count++ {
		if utf8.RuneStart(s[bytepos+count+1]) {
			return bytepos + count + 1
		}
	}
	return -1
}
