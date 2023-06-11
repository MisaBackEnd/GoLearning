package main

import "fmt"

func main() {
	sequence := []int{1, 2, 3, 4}
	res := []int{1}
	// number before
	acum := 1
	for i := 1; i < len(sequence); i++ {
		acum *= sequence[i-1]
		res = append(res, acum)
	}
	acum = 1
	for d := len(sequence) - 1; d >= 0; d-- {
		res[d] *= acum
		acum *= sequence[d]
		fmt.Println(acum)
	}
}
