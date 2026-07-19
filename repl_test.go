package main

import (
	"testing"
	"reflect"
)
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
		input: "  hello  world  ",
		expected: []string{"hello", "world"},
		},
		{
			input: "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input: " Help, Me, Please   ",
			expected: []string{"help,", "me,", "please"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual and expected does not match")
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if !reflect.DeepEqual(expectedWord, word) {
				t.Errorf("expected: %s, got: %s", expectedWord, word)
				t.Fail()
		}
		}
		
	}
}
