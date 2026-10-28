package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"}},
		{
			input:    "  foo bar  ",
			expected: []string{"foo", "bar"}},
		{
			input:    "spam   spam spam spam spam spam baked beans  spam    spam spam   ",
			expected: []string{"spam", "spam", "spam", "spam", "spam", "spam", "baked", "beans", "spam", "spam", "spam"}},
		{
			input:    "",
			expected: []string{}},
		{
			input:    "ARE YOU WINNING, SON?",
			expected: []string{"are", "you", "winning,", "son?"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		t.Logf("Expected: %v,\n Actual: %v", c.expected, actual)
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch: expected %d word(s), got %d", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Mismatch at index %d: expected %q, got %q", i, expectedWord, word)
			}
		}
	}
}
