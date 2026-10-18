package main

import "testing"

func TestLocalPart(t *testing.T) {
	cases := []struct{ in, want string }{
		{"ivan@inbox.hireme.example.com", "ivan"},
		{"ivan-2@inbox.hireme.example.com", "ivan-2"},
		{"no-at-sign", "no-at-sign"},
	}
	for _, c := range cases {
		if got := localPart(c.in); got != c.want {
			t.Errorf("localPart(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
