package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {

	testCases := []struct {
		nums    []int
		target    int
		result  []int
		msg	string
	}{
		// Example 1
		{
			nums : []int {2,7,11,15},
			target: 9,
			result: []int {0,1},
			msg : "example 1",
		},
		// Example 2
		{
			nums : []int {3,2,4},
			target: 6,
			result: []int {1,2},
			msg : "example 2",
		},
		// Example 3
		{
			nums : []int {3,3},
			target: 6,
			result: []int {0,1},
			msg : "example 3",
		},
		// Example 4
		{
			nums : []int {3,2,3},
			target: 6,
			result: []int {0,2},
			msg : "example 3",
		},
		// If number is not 
		{
			nums : []int {0,3,1,2,3},
			target: 6,
			result: []int {1,4},
			msg : "example 5",
		},
	}

	for _, testCase := range testCases {
		e := twoSum(testCase.nums, testCase.target)
		assert.Equal(t, e, testCase.result, testCase.msg)
	}

}
