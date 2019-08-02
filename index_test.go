package runes

import "testing"

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
			t.Error(n, test.Input)
		}
	}
}

func TestNext(t *testing.T) {
	type i struct {
		S string
		P int
	}
	tests := []struct {
		Input i
		Must  int
	}{
		{
			Input: i{
				S: "abcdef",
				P: 2,
			},
			Must: 3,
		},
	}

	for _, test := range tests {
		if n := Next(test.Input.S, test.Input.P); n != test.Must {
			t.Error(n, test.Input)
		}
	}
}
