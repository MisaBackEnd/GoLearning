package main

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		_, exist := numMap[num]
		if exist {
			return true
		}
		numMap[num] = i
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	containsDuplicate(nums)
}
