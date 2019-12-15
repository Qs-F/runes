package runes

import "testing"

func TestBefore(t *testing.T) {
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
			Must: 1,
		},
		{
			Input: i{
				S: "abc",
				P: 3,
			},
			Must: 2,
		},
		{
			Input: i{
				S: "abc",
				P: 0,
			},
			Must: 0,
		},
		{
			Input: i{
				S: "あいう",
				P: 5,
			},
			Must: 3,
		},
		{
			Input: i{
				S: "あい",
				P: 6,
			},
			Must: 3,
		},
		{
			Input: i{
				S: "あ",
				P: 0,
			},
			Must: 0,
		},
	}

	for _, test := range tests {
		if n := Before(test.Input.S, test.Input.P); n != test.Must {
			t.Errorf("expect %d but get %d for %s[%d]\n", test.Must, n, test.Input.S, test.Input.P)
		} else {
			t.Logf("get %d for %s[%d]\n", n, test.Input.S, test.Input.P)
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
				S: "abc",
				P: 2,
			},
			Must: 2,
		},
		{
			Input: i{
				S: "あいう",
				P: 5,
			},
			Must: 6,
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
			t.Errorf("expect %d but get %d for %s[%d]\n", test.Must, n, test.Input.S, test.Input.P)
		} else {
			t.Logf("get %d for %s[%d]\n", n, test.Input.S, test.Input.P)
		}
	}
}
