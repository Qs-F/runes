package runes

import "unicode/utf8"

const MAX_BYTE_UTF8 = 6 // 6 is the max value of utf8 byte

// Before returns the byte start position of before valid rune.
// If the before valid rune is not found, returns given bytepos.
func Before(s string, bytepos int) int {
	// check bytepos exceeds max possible length
	if bytepos > len(s)-1+MAX_BYTE_UTF8 {
		return bytepos
	}

	max := bytepos - 1
	if max > len(s)-1 {
		max = len(s) - 1
	}

	for cursor := max; cursor > bytepos-1-MAX_BYTE_UTF8; cursor-- {
		if utf8.RuneStart(s[cursor]) {
			return cursor
		}
	}
	return bytepos
}

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
