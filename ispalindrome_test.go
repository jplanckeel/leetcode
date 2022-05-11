package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {

	testCases := []struct {
		input    int
		result  bool
		msg	string
	}{
		// Example 1
		{
			input : 121,
			result: true,
			msg : "example 1",
		},
		// Example 2
		{
			input : -121,
			result: false,
			msg : "example 2",
		},
		// Example 3
		{
			input : 10,
			result: false,
			msg : "example 3",
		},
	}

	for _, testCase := range testCases {
		e := isPalindrome(testCase.input)
		assert.Equal(t, e, testCase.result, testCase.msg)
	}

}
