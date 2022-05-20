package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInteger(t *testing.T) {

	testCases := []struct {
		input    string
		result  int
		msg	string
	}{
		// Example 1
		{
			input : "III",
			result: 3,
			msg : "example 1",
		},
		// Example 2
		{
			input : "LVIII",
			result: 58,
			msg : "example 2",
		},
		// Example 3
		{
			input : "MCMXCIV",
			result: 1994,
			msg : "example 3",
		},
		// I before X
		{
			input : "IX",
			result: 9,
			msg : "I before X",
		},
		// C
		{
			input : "C",
			result: 100,
			msg : "C",
		},
		// C before D
		{
			input : "CD",
			result: 400,
			msg : "C before D",
		},
		// C before M
		{
			input : "CM",
			result: 900,
			msg : "C before M",
		},
		// M
		{
			input : "M",
			result: 1000,
			msg : "M",
		},
		{
			input : "I",
			result: 1,
			msg : "I",
		},
		{
			input : "II",
			result: 2,
			msg : "II",
		},
		{
			input : "III",
			result: 3,
			msg : "III",
		},
	}

	for _, testCase := range testCases {
		e := romanToInt(testCase.input)
		assert.Equal(t, e, testCase.result, testCase.msg)
	}

}
