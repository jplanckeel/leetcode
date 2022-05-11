package main

import "strconv"

func isPalindrome(x int) bool {

	numStr := strconv.Itoa(x)
	max := len(numStr)

	var lr []int
	var rl []int

	for _, v := range numStr {
		dlr, _ := strconv.Atoi(string(v))
		lr = append(lr, dlr)
		drl, _ := strconv.Atoi(string(numStr[max]))
		rl = append(lr, drl)
		max -= 1
	}


    for i := range rl {
        if rl[i] != lr[i] {
            return false
        }
    }
    return true

}
