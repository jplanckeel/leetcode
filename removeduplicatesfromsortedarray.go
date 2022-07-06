package main

func removeDuplicates(nums []int) (int,[]int) {
    if len(nums) <= 1 {
        return len(nums), nums
    }
	
	var i, j int = 0, 1

    for j <  len(nums) {
		if nums[i] == nums[j]{
			j++
		} else {
			nums[i+1] = nums[j]
			j++
			i++
		} 
		
	}
	k := i
	for i <  len(nums) -1 {
		nums[i+1] = 0
			i++
	} 

	return k+1, nums
}
