package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	len_text := len(text)
	for i := len_text - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if strings.ToLower(textReverse) == strings.ToLower(text) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	// arrays
	var arr [5]int
	arr[0] = 1
	arr[2] = 2
	fmt.Println(arr)
	arr2 := [...]string{"uno", "dos"}
	fmt.Println(arr2)
	arr3 := [2]string{"tres", "cuatro"}
	fmt.Println(arr3)

	// slices
	var letters []string
	letters = append(letters, "a")
	letters = append(letters, "b")
	fmt.Println(letters)

	// iterando en slices
	frase := []string{"hola", "que", "hace?"}
	for i, val := range frase {
		fmt.Println(i, val)
	}
	text := "hola"
	fmt.Println(text[0])
	isPalindromo("Amor a Rom")

	// maps o hash tables (implementan concurrencia de forma nativa)
	personAges := make(map[string]int)
	personAges["Misael"] = 34
	personAges["Katherine"] = 34
	personAges["Ines"] = 32
	personAges["Daniel"] = 32
	fmt.Println(personAges)

	//Recorrer un map
	for i, v := range personAges {
		fmt.Println(i, v)
	}

	// acceder a un valor
	value, ok := personAges["Misaelo"]
	println(value, ok)
}
