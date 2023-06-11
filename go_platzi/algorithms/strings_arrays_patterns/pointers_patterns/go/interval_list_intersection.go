package main

import "fmt"

func get_intersection(first []int, second []int) (res []int, exist bool) {
	intersection := make([]int, 2)

	if second[0] >= first[0] && second[0] <= first[1] {
		intersection[0] = second[0]
		if second[1] >= first[1] {
			intersection[1] = first[1]
		} else {
			intersection[1] = second[1]
		}
		return intersection, true
	}
	if first[0] >= second[0] && first[0] <= second[1] {
		intersection[0] = first[0]
		if second[1] >= first[1] {
			intersection[1] = first[1]
		} else {
			intersection[1] = second[1]
		}
		return intersection, true
	}
	return intersection, false
}

func main() {
	// one := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	// two := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	one := [][]int{{1, 3}, {5, 9}}
	two := [][]int{}
	i := 0
	d := 0
	result := [][]int{}
	res := []int{}
	var intersect bool
	for d < len(two) && i < len(one) {
		res, intersect = get_intersection(one[i], two[d])
		if intersect {
			result = append(result, res)
		}
		if one[i][1] >= two[d][1] {
			d++
		} else {
			i++
		}
	}
	fmt.Println(result)

}
