package runes

import "testing"

func TestCount(t *testing.T) {
	type Input struct {
		S string
		N int
	}
	tests := []struct {
		Input Input
		Must  int
	}{
		{
			Input: Input{
				S: "あいう",
				N: -1,
			},
			Must: 3,
		},
		{
			Input: Input{
				S: "あ",
				N: 2,
			},
			Must: 1,
		},
		{
			Input: Input{
				S: "あ",
				N: 0,
			},
			Must: 0,
		},
		{
			Input: Input{
				S: "あ",
				N: 1,
			},
			Must: 0,
		},
		{
			Input: Input{
				S: "あ",
				N: 2,
			},
			Must: 0,
		},
		{
			Input: Input{
				S: "あい",
				N: 4,
			},
			Must: 1,
		},
	}

	for _, test := range tests {
		c := Count(test.Input.S, test.Input.N)
		if c != test.Must {
			t.Errorf("expect: %d but get %d for %s\n", test.Must, c, test.Input.S)
		} else {
			t.Logf("get %d for %s\n", c, test.Input.S)
		}
	}
}

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
		{
			Input: i{
				S: "あいうえお日本語",
				P: 3,
			},
			Must: 6,
		},
	}

	for _, test := range tests {
		if n := Next(test.Input.S, test.Input.P); n != test.Must {
			t.Error(n, test.Input)
		}
	}
}
