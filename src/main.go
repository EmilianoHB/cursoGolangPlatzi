package main

import "fmt"

func main() {
	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	//Division
	result = y / x
	fmt.Println("Division: ", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Retos
	//Calcular el area de un rectangulo, trapecio y de un circulo
	//Area rectangulo
	rLargo := 6
	rAncho := 2
	AreaRectangulo := rLargo * rAncho
	fmt.Println("El area del rectangulo es: ", AreaRectangulo)

	//Area Trapecio
	base1 := 10
	base2 := 4
	alturaT := 4
	AreaTrapecio := ((base1 + base2) * alturaT) / 2
	fmt.Println("El area del trapecio es: ", AreaTrapecio)

	//Area Circulo
	radio := 0.5
	const pi = 3.14
	AreaCirculo := pi * (radio * radio)
	fmt.Println("El area del circulo es: ", AreaCirculo)
}
