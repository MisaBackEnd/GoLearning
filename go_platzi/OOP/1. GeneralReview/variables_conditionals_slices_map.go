package main

import (
	"fmt"
	"strconv"
)

func main() {
	// variables
	var x int // declaration
	x = 8     // assignation
	y := 8    // declaration and asignation
	fmt.Println(x, y)

	// no explicit error handling
	myValue, err := strconv.ParseInt("NaN", 0, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myValue)
	}

}
