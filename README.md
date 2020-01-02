# pkg `runes`

Go pkg inspried by `strings` package, for `runes`.

Although `strings` and `unicode/utf8` pkg is great, but some of them are behaving weirdly.

For example,

- If the byte slice is incomplete, then `utf8.CountXXX` counts the length of bytes and add them.

- `strings.Index` returns `0` if passing empty string `""` to its parameter `substr`.

With this pkg, these points are fixed.

![Test](https://github.com/Qs-F/runes/workflows/test/badge.svg)
[![GoDoc](https://godoc.org/github.com/Qs-F/runes?status.svg)](https://godoc.org/github.com/Qs-F/runes)
[![Go Report Card](https://goreportcard.com/badge/github.com/Qs-F/runes)](https://goreportcard.com/report/github.com/Qs-F/runes)

## Installation

`go get github.com/Qs-F/runes`

## Usage

See [GoDoc](https://godoc.org/github.com/Qs-F/runes)

## License

MIT
