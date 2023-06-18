package main

import "fmt"

func maxNumberOfOnes(nums []int) int {
	m1 := 0
	c := 0
	for _, n := range nums {
		if n == 0 {
			c = 0
		} else {
			c++
			if c > m1 {
				m1 = c
			}
		}
	}
	return m1
}

func main() {
	numbers := []int{
		0, 1, 0, 1, 1, 0,
	}
	result := maxNumberOfOnes(numbers)
	fmt.Println(result)
}
