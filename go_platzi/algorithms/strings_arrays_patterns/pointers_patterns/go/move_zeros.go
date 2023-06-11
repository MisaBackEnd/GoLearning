package main

func moveZeroes(nums []int) {
	lastNonZero := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	i := -1
	d := 0
	for d < len(nums) {
		if nums[d] == 0 {
			if i == -1 {
				i = d
			}
		} else {
			if i >= 0 {
				nums[i], nums[d] = nums[d], nums[i]
				i++
			}
		}
		d++

	}
}
