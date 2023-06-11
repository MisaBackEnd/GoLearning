package main

import "fmt"

func main() {
	s := []string{"h", "e", "l", "l", "o"}
	for i, _ := range s[:len(s)/2] {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	fmt.Println(s)
}
