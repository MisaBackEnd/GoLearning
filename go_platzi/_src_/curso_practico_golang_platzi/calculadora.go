package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

// receiver function
func (calc) operate(entrada string, operador string) int {
	valores := strings.Split(entrada, operador)
	op1 := parsear(valores[0])
	op2 := parsear(valores[1])
	switch operador {
	case "+":
		return op1 + op2
	case "-":
		return op1 - op2
	case "*":
		return op1 * op2
	case "/":
		return op1 / op2
	default:
		return 0
	}
}

func parsear(entrada string) int {
	op, _ := strconv.Atoi(entrada)
	return op
}

func leerEntrada() string {
	fmt.Println("Bienvenido! Para utilizar la calculadora debes escribir los dos operadores")
	fmt.Println("y el simbolo de la operacion que desees realizar. Ejemplo 2+2.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func leerOperador() string {
	fmt.Println("Suministra el tipo de operacion a realizar +-*/")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entrada := leerEntrada()
	operador := leerOperador()
	calculadora := calc{}
	result := calculadora.operate(entrada, operador)
	fmt.Println(result)
}
