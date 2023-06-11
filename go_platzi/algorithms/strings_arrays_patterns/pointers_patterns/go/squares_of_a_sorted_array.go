package main

import "fmt"

func sortedSquares(nums []int) []int {
	i, d, p := 0, len(nums)-1, len(nums)-1
	result := make([]int, len(nums))
	for d >= i {
		if nums[i]*nums[i] >= nums[d]*nums[d] {
			result[p] = nums[i] * nums[i]
			i++
		} else {
			result[p] = nums[d] * nums[d]
			d--
		}
		p--
	}
	return result
}

func main() {
	// fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	// fmt.Println(sortedSquares([]int{-5, -3, -2, -1}))
	// fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
	fmt.Println(sortedSquares([]int{-10000, -9999, -7, -5, 0, 0, 10000}))
}
