package main

import (
	"fmt"
	"strings"
)

func isAlienSorted(words []string, alphabet string) bool {
	// convert alphabet to hash table
	alphabetHashTable := make(map[string]int)
	for i, l := range alphabet {
		alphabetHashTable[string(l)] = i
	}
	a, exist := alphabetHashTable[string('b')]
	fmt.Println(a, exist)
	// iterate words
	p := 1
	j := 0

	for p < len(words) {
		j = 0
		for j <= len(words[p]) {
			if len(words[p-1]) > len(words[p]) && len(words[p]) == j {
				return false
			}
			if j < len(words[p]) && j < len(words[p-1]) && words[p][j] != words[p-1][j] {
				if alphabetHashTable[string(words[p][j])] < alphabetHashTable[string(words[p-1][j])] {
					return false
				}
				j = len(words[p]) + 1
			}
			j++
		}
		p++
	}
	return true
}

func isAlienSortedOptimized(words []string, order string) bool {
	orderMap := map[string]int{}
	var word []string
	var nextWord []string
	for key, value := range strings.Split(order, "") {
		orderMap[value] = key
	}
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			word = strings.Split(words[i], "")
			nextWord = strings.Split(words[i+1], "")

			if j >= len(words[i+1]) {
				return false
			}

			if orderMap[word[j]] != orderMap[nextWord[j]] {
				if orderMap[word[j]] > orderMap[nextWord[j]] {
					return false
				}
				break
			}
		}
	}
	return true
}

func main() {
	words := []string{"apple", "app"}
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	result1 := isAlienSorted(words, alphabet)
	fmt.Println("Result 1:", result1)
	words = []string{"hello", "leetcode"}
	alphabet = "hlabcdefgijkmnopqrstuvwxyz"
	result2 := isAlienSorted(words, alphabet)
	fmt.Println("result2", result2)
	words = []string{"kuvp", "q"}
	alphabet = "ngxlkthsjuoqcpavbfdermiywz"
	result3 := isAlienSorted(words, alphabet)
	fmt.Println("result3", result3)
	words = []string{"hello", "hello"}
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	result4 := isAlienSorted(words, alphabet)
	fmt.Println("result4", result4)
}
