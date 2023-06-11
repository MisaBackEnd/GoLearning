package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	l, r, lmh, rmh, w := 0, len(height)-1, height[0], height[len(height)-1], 0
	for l < r {
		// the most probable thing is that if i encounter a
		// bigger height in l water is gonna be trapped.
		if height[l] < height[r] {
			if height[l] < lmh {
				w = w + lmh - height[l]
			} else {
				lmh = height[l]
			}
			l++
		} else {
			if height[r] < rmh {
				w = w + rmh - height[r]
			} else {
				rmh = height[r]
			}
			r--
		}
	}
	fmt.Println(w)
}
