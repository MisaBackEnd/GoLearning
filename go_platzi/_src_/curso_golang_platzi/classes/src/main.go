package main

import (
	pk "classes/src/mypackage"
	"fmt"
)

func main() {
	var mycar pk.CarPublic
	mycar.Brand = "Chevrolet Spark"
	mycar.Year = 2011
	fmt.Println(mycar)
}
