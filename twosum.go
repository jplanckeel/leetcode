package main

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := range nums {
		t := i + 1
		for t < l {
			if nums[i] + nums[t] == target {
				var r []int
				r = append(r,i,t)
				return r
			}
			t += 1
		}
	}
	return nil
}
