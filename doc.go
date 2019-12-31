// pkg `runes` is the golang lib package for runes, which is inspired by standard lib `strings`.
//
// Although `strings` and `unicode/utf8` pkg is great, but some of them are behaving weirdly.
//
// For example,
//
// - If the byte slice is incomplete, then `utf8.CountXXX` counts the length of bytes and add them.
//
// - `strings.Index` returns `0` if passing empty string `""` to its parameter `substr`.
//
// With this pkg, these points are fixed.
package runes
