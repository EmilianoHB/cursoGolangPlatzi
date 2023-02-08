package main

import (
	"fmt"
	"log"
	"strconv"
)

func parOimpar(a int) {
	if a%2 == 0 {
		par := "Es par"
		fmt.Println(par)
	} else {
		impar := "Es impar"
		fmt.Println(impar)
	}
}

func iniciarsesion(a string, b int) {
	if a == "Emiliano" && b == 1234 {
		fmt.Println("Correcto")
	} else {
		fmt.Println("Incorrecto")
	}

}
func main() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//Whith and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	//Whit or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Se cumple la condicion")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

	//Retos
	parOimpar(9)
	iniciarsesion("Emiliano", 1234)
}
