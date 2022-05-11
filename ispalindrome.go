package main

import "strconv"

func isPalindrome(x int) bool {

	numStr := strconv.Itoa(x)
	var l []string

	for _, v := range numStr {
		l = append(l, string(v))
	}


	max := len(numStr)


    for i := range l {
		max += -1
        if l[i] != l[max] {
            return false
        }
    }
    return true

}
