package main

import "testing"

type testSuite struct {
	arg, expected string
}

var testSuits = []testSuite{
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd", "abcd"},
	{"3abcd", ""},
	{"45", ""},
	{"aaaa10b", ""},
	{"aaa0b", "aab"},
	{"d\n5abc", "d\n\n\n\n\nabc"},
	{`qwe\4\5`, "qwe45"},
	{`qwe\45`, "qwe44444"},
	{`qwe\\5`, `qwe\\\\\`},
	{`qw\ne`, ""},
}

func TestUnpack(t *testing.T) {
	for _, test := range testSuits {
		if output := unpack(test.arg); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		} else {
			t.Logf("Success %q is equal %q", output, test.expected)
		}
	}
}
