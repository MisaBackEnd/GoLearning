package main

// Fuck yeah the most optimal solution!
func removeElement(nums []int, val int) int {
	p := 0
	for _, n := range nums {
		if n != val {
			nums[p] = n
			p++
		}
	}
	return p
}
