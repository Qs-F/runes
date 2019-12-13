package runes

import "unicode/utf8"

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
