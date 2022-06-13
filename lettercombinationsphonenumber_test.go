package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {

	testCases := []struct {
		input    string
		result  []string
		msg	string
	}{	
		{
			input : "234",
			result : []string {"adg","adh","adi","aeg","aeh","aei","afg","afh","afi","bdg","bdh","bdi","beg","beh","bei","bfg","bfh","bfi","cdg","cdh","cdi","ceg","ceh","cei","cfg","cfh","cfi"},
			msg : "submit result",
		},
		{
			input : "23",
			result : []string {"ad","ae","af","bd","be","bf","cd","ce","cf"},
			msg : "example 1",
		},
		{
			input : "",
			result : []string {},
			msg : "example 2",
		},
		{
			input : "2",
			result : []string {"a","b","c"},
			msg : "example 3",
		},

	}
	
	for _, testCase := range testCases {
		e := letterCombinations(testCase.input)
		assert.Equal(t, e, testCase.result, testCase.msg)
	}

}
