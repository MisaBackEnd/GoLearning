package main

import "fmt"

func mergeTwoSortedSlices(nums1 *[]int, m int, nums2 *[]int, n int) {
	p1 := m - 1
	p2 := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if p2 < 0 {
			break
		}
		if p1 >= 0 && (*nums1)[p1] >= (*nums2)[p2] {
			(*nums1)[i] = (*nums1)[p1]
			p1--
		}
		if p1 < 0 || (p1 >= 0 && (*nums1)[p1] < (*nums2)[p2]) {
			(*nums1)[i] = (*nums2)[p2]
			p2--
		}
	}
}

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	m := 1
	n := 1
	mergeTwoSortedSlices(&nums1, m, &nums2, n)
	fmt.Println(nums1)
}
