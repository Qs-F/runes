package runes

import (
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		Input []string
		Must  int
	}{
		{
			Input: []string{
				"abcdefghijk",
				"def",
			},
			Must: 3,
		},
		{
			Input: []string{
				"abc日本語",
				"本語",
			},
			Must: 4,
		},
		{
			Input: []string{
				"abcdefghijk",
				"xyz",
			},
			Must: -1,
		},
		{
			Input: []string{
				"abcdefghijk",
				"",
			},
			Must: -1,
		},
		{
			Input: []string{
				"",
				"abc",
			},
			Must: -1,
		},
	}

	for _, test := range tests {
		if n := Index(test.Input[0], test.Input[1]); n != test.Must {
			t.Errorf("expect %d but get %d for '%s' of '%s'\n", test.Must, n, test.Input[0], test.Input[1])
		} else {
			t.Logf("get %d for '%s' of '%s'\n", test.Must, test.Input[0], test.Input[1])
		}
	}
}

func TestIndexAll(t *testing.T) {
	tests := []struct {
		Input []string
		Must  []int
	}{
		{
			Input: []string{
				"abcabcabc",
				"abc",
			},
			Must: []int{0, 3, 6},
		},
		{
			Input: []string{
				"aaaaaaaa",
				"a",
			},
			Must: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			Input: []string{
				"abcabcabc",
				"",
			},
			Must: []int{},
		},
		{
			Input: []string{
				"",
				"abc",
			},
			Must: []int{},
		},
		{
			Input: []string{
				"ああいうえおabcdあ",
				"あ",
			},
			Must: []int{0, 1, 10},
		},
		{
			Input: []string{
				"ああいうえおabcdあいう",
				"あ",
			},
			Must: []int{0, 1, 10},
		},
	}

	for _, test := range tests {
		if n := IndexAll(test.Input[0], test.Input[1]); !reflect.DeepEqual(n, test.Must) {
			t.Errorf("expect %d but get %d for '%s' of '%s'\n", test.Must, n, test.Input[0], test.Input[1])
		} else {
			t.Logf("get %d for '%s' of '%s'\n", test.Must, test.Input[0], test.Input[1])
		}
	}
}
