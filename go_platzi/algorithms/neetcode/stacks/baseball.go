package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	stack := []int{}
	for _, operation := range operations {
		switch operation {
		case "+":
			newScore := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, newScore)
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			stack = stack[:len(stack)-1]
		default:
			value, err := strconv.Atoi(operation)
			if err == nil {
				stack = append(stack, value)
			}
		}
	}
	result := 0
	for _, item := range stack {
		result += item
	}
	return result
}

func main() {
	result_all := calPoints([]string{"5", "2", "C", "D", "+"})
	fmt.Println(result_all)
}
