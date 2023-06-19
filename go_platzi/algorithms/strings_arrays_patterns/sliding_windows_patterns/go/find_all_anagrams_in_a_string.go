package main

import "reflect"

func findAnagrams(s string, p string) []int {
	// create a map for letter frequencies
	// in p string.
	pm := map[string]int{}
	result := []int{}
	for _, l := range p {
		_, exist := pm[string(l)]
		if !exist {
			pm[string(l)] = 1
		} else {
			pm[string(l)] += 1
		}
	}

	// initialize map for letter frequencies in current window
	cwm := make(map[string]int, len(pm))
	start := -1

	for i, l := range s {
		// fill the current window dictionary with
		// the corresponding letter frequencies
		_, exist := cwm[string(l)]
		if !exist {
			cwm[string(l)] = 1
		} else {
			cwm[string(l)] += 1
		}
		// check that the current window has the same length as p
		if i-start == len(p) {
			anagram := reflect.DeepEqual(cwm, pm)
			si := start + 1     // real start index for the window
			sc := string(s[si]) // start character of the window
			if anagram {
				result = append(result, si)
			}
			// descrease frequency of the caracter in the last window
			cwm[sc] -= 1
			// deleted if necessary
			if cwm[sc] == 0 {
				delete(cwm, sc)
			}
			start++
		}
	}
	return result
}
