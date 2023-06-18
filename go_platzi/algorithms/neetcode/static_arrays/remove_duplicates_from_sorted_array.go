package main

func removeDuplicates(nums []int) int {
	pr := 0 // pointer that has the position where the next non repeating value should be placed
	r_ht := make(map[int]int, len(nums))
	for i, n := range nums {
		_, exist := r_ht[n]
		if !exist {
			r_ht[n] = i
			nums[pr] = n
			pr++
		}
	}
	return pr + 1
}

func optimalSolution(nums []int) int {
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
