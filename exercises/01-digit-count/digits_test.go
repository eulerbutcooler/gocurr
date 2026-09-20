package main

import "testing"

func TestCountDigitsInWords(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		sent string
		want int
	}{
		{"example", "meo1 mai2 hu catik21123a", 7},
		{"no digits", "no digits here", 0},
		{"all digits", "123 456", 6},
		{"empty", "", 0},
		{"single word", "abc42", 2},
		{"mixed", "a1b2c3", 3},
		{"repeated words", "a1 a1 a1", 3},
		{"arabic-indic", "١٢٣", 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			for _, im := range impls {
				if got := im.fn(tc.sent); got != tc.want {
					t.Errorf("%s(%q) = %d, want %d", im.name, tc.sent, got, tc.want)
				}
			}
		})
	}
}

func TestCountDigits(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		in   string
		want int
	}{
		{"empty", "", 0},
		{"letters only", "abc", 0},
		{"one digit", "a1", 1},
		{"run", "catik21123a", 5},
		{"emoji", "😭😭😭😭😭1", 1},
		{"arabic-indic", "١٢٣", 3},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := countDigits(tc.in); got != tc.want {
				t.Errorf("countDigits(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
