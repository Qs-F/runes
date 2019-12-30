package runes

import (
	"unicode/utf8"
)

// Count returns the number of str in rune.
// If bytecnt >= 0, then counts str[:bytescnt].
// If bytecnt < 0 (mostly -1), then counts all the str.
func Count(str string, bytecnt int) int {
	if bytecnt < 0 {
		bytecnt = len(str)
	}

	if utf8.ValidString(str[Before(str, bytecnt):bytecnt]) {
		return utf8.RuneCountInString(str[:bytecnt])
	}

	l := len(str[:bytecnt])
	offset := l - Before(str, bytecnt) - 1
	return utf8.RuneCountInString(str[:bytecnt]) - offset
}
