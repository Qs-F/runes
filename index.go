package runes

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Index returns index number in rune
func Index(s, substr string) int {
	n := strings.Index(s, substr)
	if n < 0 {
		return n
	}
	return Count(s, n)
}

// Count returns the number of str in rune.
// If bytecnt >= 0, then counts str[:bytescnt].
// If bytecnt < 0 (mostly -1), then counts all the str.
func Count(str string, bytecnt int) int {
	if bytecnt < 0 {
		return utf8.RuneCountInString(str)
	}
	fmt.Println([]byte(str)[:bytecnt])
	return utf8.RuneCountInString(str[:bytecnt])
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
