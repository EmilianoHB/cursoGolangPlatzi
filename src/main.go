package main

import (
	"fmt"
)

func main() {
	//Declacracion de constatntes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Derclaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float32
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	areaCuadrado := base * altura / 2
	fmt.Println(areaCuadrado)
}
