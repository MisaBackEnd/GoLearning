package main

import (
	"fmt"
	"log"
	"strconv"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func returnValue(a, b int) int {
	return a * b
}

func main() {
	// declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	// zero values, go asigna un valor default
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado:", areaCuadrado)

	// suma
	x := 2
	y := 1
	result := x + y
	fmt.Println(result)

	// resta
	result = y - x
	fmt.Println(result)

	// Multiplicacion
	result = 2 * 1
	fmt.Println(result)

	// division
	result = 5 / 2
	fmt.Println(result)

	// Modulo
	result = 2 % 1
	fmt.Println(result)

	// incrementar
	x++
	fmt.Println(x)
	// decrementar
	x--
	fmt.Println(x)

	// Tipos de datos
	// int depende del os SI uint es para numeros positivos
	// int8 -> 8 bits -128 a 127
	// int16 -> 16 bits -2^15 a 2^15 -1
	// int32 -> 32 bits -2^31 a 2^31 -1
	// int64 = 64bits = -2^63 a 2^63-1

	// float32 y float64
	// textos "" y booleanos true o false

	// Complex64 y Complex128

	//fmt
	yo := "Misael"
	precio_mac := 16
	fmt.Printf("%s va a comprar un macbook pro de %d millones de pesos\n", yo, precio_mac)
	fmt.Printf("%v va a comprar un macbook pro de %v millones de pesos\n", yo, precio_mac)
	// el tipo de dato
	fmt.Printf("Tipo de dato para precio mac: %T\n", precio_mac)

	// funciones y funciones anonimas
	normalFunction("hola mundo")
	returnedValue := returnValue(1, 23)
	fmt.Println(returnedValue)

	// loops
	// for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	// for forever
	// for
	// 	lo que sea
	//
	result = 5 / 2
	result2 := 5 / 2.0
	fmt.Println(result)
	fmt.Println(result2)

	// Parsing
	value, err := strconv.Atoi("5")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr, len(arr), cap(arr))

}
