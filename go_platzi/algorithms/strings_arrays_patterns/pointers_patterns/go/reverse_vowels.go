package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	v := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	vr := strings.Split(s, "")
	i, d := 0, len(s)-1
	for i < d {
		_, exist_i := v[string(s[i])]
		_, exist_d := v[string(s[d])]
		switch {
		case exist_d && exist_i:
			vr[i], vr[d] = vr[d], vr[i]
			i++
			d--
		case exist_d && !exist_i:
			i++
		case !exist_d && exist_i:
			d--
		default:
			d--
			i++
		}
	}

	return strings.Join(vr, "")
}

func main() {
	fmt.Println(reverseVowels("hellow"))
}
