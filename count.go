package runes

import (
	"unicode/utf8"
)

// Count returns the number of str in rune.
// If bytecnt >= 0, then counts str[:bytescnt].
// If bytecnt < 0 (mostly -1), then counts all the str.
func Count(str string, bytecnt int) int {
	if bytecnt < 0 {
		return utf8.RuneCountInString(str)
	}
	l := len(str[:bytecnt])
	offset := l - Before(str, bytecnt)
	return utf8.RuneCountInString(str[:bytecnt]) - offset
}
