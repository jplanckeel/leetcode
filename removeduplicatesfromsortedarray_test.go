package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestRemoveDuplicates(t *testing.T) {

	testCases := []struct {
		input    []int
		expectedNums []int
		result  int
		msg	string
	}{
		// Example 1
		{
			input : []int{1,1,2},
			expectedNums:[]int{1,2,0},
			result: 2,
			msg : "example 1",
		},
		// Example 2
		{
			input : []int{0,0,1,1,1,2,2,3,3,4},
			expectedNums:[]int{0,1,2,3,4,0,0,0,0,0},
			result: 5,
			msg : "example 2",
		},
		// Negative Number
		{
			input : []int{-2,-2,1,1,1,2,2,3,3,4},
			expectedNums:[]int{-2,1,2,3,4,0,0,0,0,0},
			result: 5,
			msg : "With negative number",
		},
		// No duplicate
		{
			input : []int{-2,1,2,3,4},
			expectedNums:[]int{-2,1,2,3,4},
			result: 5,
			msg : "No duplicate",
		},
	}

	for _, testCase := range testCases {
		e, numsE := removeDuplicates(testCase.input)
		assert.Equal(t, e, testCase.result, testCase.msg)
		assert.Equal(t, numsE, testCase.expectedNums, testCase.msg)
	}

}
