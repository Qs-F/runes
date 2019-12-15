package runes

import "unicode/utf8"

const MAX_BYTE_UTF8 = 6 // 6 is the max value of utf8 byte

// Next returns the byte start position of next valid rune.
// If the next valid rune is not found, returns given bytepos.
func Next(s string, bytepos int) int {
	// check bytepos exceeds max possible length
	if bytepos > len(s)-1 {
		return bytepos
	}

	max := bytepos + 1 + MAX_BYTE_UTF8
	if max > len(s)-1 {
		max = len(s) - 1
	}

	for cursor := bytepos + 1; cursor < max; cursor++ { // count starts from 1, cuz this is Next()
		if utf8.RuneStart(s[cursor]) {
			return cursor
		}
	}
	return bytepos
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
