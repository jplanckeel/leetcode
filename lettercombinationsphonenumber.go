package main

var dict = map[byte][]string {
    '2' : { "a", "b", "c"},
    '3' : { "d", "e", "f"},
    '4' : { "g", "h", "i"},
    '5' : { "j", "k", "l"},
    '6' : { "m", "n", "o"},
    '7' : { "p", "q", "r", "s"},
    '8' : { "t", "u", "v"},
    '9' : { "w", "x", "y", "z"},

}




func letterCombinations(digits string) []string {

    if digits == "" {
        return []string{}
    }
       
    var max int = 1
    for i := range digits {
        max = max * len(dict[digits[i]])
    }

    result := make([]string, max)

	count := max
	for i := range digits {
		// count means that a char would repeat count times
		count = count / len(dict[digits[i]])
		for j := range result {
			// to rolling the chars slice
			result[j] = result[j] + dict[digits[i]][(j/count)%len(dict[digits[i]])]
		}
	}
    

    return result
}
