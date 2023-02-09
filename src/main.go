package main

import "fmt"

type auto struct {
	marca    string
	modelo   int
	color    string
	puertas  int
	asientos int
}

func main() {
	miAuto := auto{
		marca:    "Chevrolet",
		modelo:   2019,
		color:    "Negro",
		puertas:  5,
		asientos: 5,
	}
	fmt.Println(miAuto)

	//Otra forma
	var otroAuto auto
	otroAuto.marca = "Ford"
	otroAuto.modelo = 2012
	otroAuto.color = "Rojo"
	otroAuto.puertas = 3
	otroAuto.asientos = 2

	fmt.Println(otroAuto)
}
