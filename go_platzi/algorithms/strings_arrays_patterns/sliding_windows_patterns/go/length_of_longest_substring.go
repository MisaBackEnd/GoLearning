package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	dict := make(map[string]int)
	maxLength, ip := 0, 0

	for i, l := range s {
		letter := string(l)
		dp, exists := dict[letter]
		// revisar si existe y si debe ser actualizado al index actual
		// en el diccionario.
		if !exists {
			dict[letter] = i
		} else {
			dict[letter] = i
			// fmt.Println("-------------------------")
			// fmt.Println("i", i)
			// fmt.Println("ip", ip)
			// fmt.Println("letter", letter)
			// fmt.Println("i-ip", i-ip)
			switch {
			case dp > ip:
				if dp < i {
					ip = dp + 1
				} else {
					ip = i
				}
			case dp == ip:
				ip++
			default:
				// nothing happens
			}
		}
		fmt.Printf("final i%d - ip%d = %d\n", i, ip, i-ip)
		maxLength = max(maxLength, len(s[ip:i+1]))

		// actualizar el pointer del inicio de la subcadena paramath
		// quede fuera del rango de repeticion.

		// existen dos formas en las que se puede mover el pointer de inicio
		// de la subcadena, se puede mover de a uno o saltar directamente al index

		// si se encuentra al inicio de la subcadena entonces se debe:
		//		actualizar el index del diccionario al indice donde se encuentra
		//		repetido.
		//		se debe correr el pointer de inicio una posicion.
		// si se encuentra en una posicion mayor al inicio de la subcadena:
		// 		actualizar el index del diccionario
		//		correr el pointer de inicio a la posicion donde se encontro la repeticion

	}
	fmt.Println(dict)
	return maxLength
}

func main() {
	st1 := "abcabcab"
	result := lengthOfLongestSubstring(st1)
	fmt.Println(result, st1)
	st1 = "bbbbb"
	result = lengthOfLongestSubstring(st1)
	fmt.Println(result, st1)
	st1 = "pwwkew"
	result = lengthOfLongestSubstring(st1)
	fmt.Println(result, st1)
	st1 = "ohvhjdml"
	result = lengthOfLongestSubstring(st1)
	fmt.Println(result, st1)
	st1 = "dvdf"
	result = lengthOfLongestSubstring(st1)
	fmt.Println(result, st1)
	st1 = "au"
	result = lengthOfLongestSubstring(st1)
	fmt.Println(result, st1)
	st1 = "aab"
	result = lengthOfLongestSubstring(st1)
	fmt.Println(result, st1)
	st1 = "tmmzuxt"
	result = lengthOfLongestSubstring(st1)
	fmt.Println(result, st1)
}
