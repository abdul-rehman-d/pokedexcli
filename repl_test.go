package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO World",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Lengths are not equal: %v vs %v", len(actual), len(c.expected))
			continue
		}
		for idx := range actual {
			if actual[idx] != c.expected[idx] {
				t.Errorf("%v does not match %v", actual[idx], c.expected[idx])
				continue
			}
		}
	}
}
