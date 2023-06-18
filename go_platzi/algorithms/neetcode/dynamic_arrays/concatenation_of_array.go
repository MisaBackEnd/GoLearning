package main

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}

// according to leetcode
func optimal(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i, v := range nums {
		ans[i] = v
		ans[i+len(nums)] = v
	}
	return ans
}
