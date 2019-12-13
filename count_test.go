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
