package runes

import "testing"

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
