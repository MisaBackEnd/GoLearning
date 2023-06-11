package main

import "fmt"

func main() {
	// keys are strings and values are int
	m1 := make(map[string]int)
	m1["a"] = 8
	fmt.Println(m1)
	fmt.Println(m1["a"])
}
