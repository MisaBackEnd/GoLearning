package main

import "fmt"

func validPalindrome(s string) bool {
	i, d, r := 0, len(s)-1, 0
	br_i, br_d := 0, len(s)-1
	for i < d {
		if s[i] != s[d] {
			r++
			i++
			if r == 1 {
				br_i, br_d = i-1, d-1
			}
			if r > 1 {
				break
			}
		} else {
			i++
			d--
		}
	}
	if r <= 1 {
		return true
	}
	r = 0
	i = br_i
	d = br_d
	for i < d {
		if s[i] != s[d] {
			r++
			d--
			if r > 1 {
				break
			}
		} else {
			i++
			d--
		}
	}
	if r < 1 {
		return true
	}
	return false
}

func main() {
	result := validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga")
	fmt.Println(result)
}
