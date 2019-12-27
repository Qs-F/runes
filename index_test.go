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
	}

	for _, test := range tests {
		if n := IndexAll(test.Input[0], test.Input[1]); !reflect.DeepEqual(n, test.Must) {
			t.Errorf("expect %d but get %d for '%s' of '%s'\n", test.Must, n, test.Input[0], test.Input[1])
		} else {
			t.Logf("get %d for '%s' of '%s'\n", test.Must, test.Input[0], test.Input[1])
		}
	}
}
