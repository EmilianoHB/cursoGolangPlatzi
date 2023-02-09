package main

import (
	pk "cursoGolangPlatzi/src/mypackage"
	"fmt"
)

func main() {
	//Forma 1
	var miAuto1 pk.AutoPublico
	miAuto1.Marca = "Ferrari"
	miAuto1.Modelo = 2023
	miAuto1.Color = "Rojo"
	miAuto1.Puertas = 3
	miAuto1.Asientos = 2
	fmt.Println(miAuto1)

	//Forma 2
	miAuto := pk.AutoPublico{
		Marca:    "Chevrolet",
		Modelo:   2019,
		Color:    "Negro",
		Puertas:  5,
		Asientos: 5,
	}
	fmt.Println(miAuto)

	pk.ImprimirMensaje("Hola como andan, ya no hace falta comentar")
}
