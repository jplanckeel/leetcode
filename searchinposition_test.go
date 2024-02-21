package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {

	testCases := []struct {
		input  []int
		target int
		output int
		msg    string
	}{
		// Example 1
		{
			input:  []int{1, 3, 5, 6},
			output: 2,
			target: 5,
			msg:    "example 1",
		},
		// Example 2
		{
			input:  []int{1, 3, 5, 6},
			output: 1,
			target: 2,
			msg:    "example 2",
		},
		// Example 3
		{
			input:  []int{1, 3, 5, 6},
			output: 4,
			target: 7,
			msg:    "example 3",
		},
	}

	for _, testCase := range testCases {
		e := searchInsert(testCase.input, testCase.target)
		assert.Equal(t, e, testCase.output, testCase.msg)
	}
}
