package main

import "fmt"

func main() {
	ba := 0
	c := 0
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	i := 0
	d := len(heights) - 1
	for d >= i {
		if heights[d] >= heights[i] {
			c = heights[i] * (d - i)
			if c > ba {
				ba = c
			}
			i++
		} else {
			c = heights[d] * (d - i)
			if c > ba {
				ba = c
			}
			d--
		}
	}
	fmt.Println(ba)
}

// most optimal solution
func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1

	result := 0
	for l <= r {
		area := (r - l) * min(height[l], height[r])
		result = max(result, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
