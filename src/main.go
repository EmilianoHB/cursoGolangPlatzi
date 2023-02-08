package main

import (
	"fmt"
	"strings"
)

func isPalidnromo(text string) {
	var textReverse string
	text = strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es un palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}
func main() {
	slice := []string{"Hola", "que", "haces?"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}
	//----------------------------------------------
	for _, valor := range slice {
		fmt.Println(valor)
	}
	//----------------------------------------------
	for i := range slice {
		fmt.Println(i)
	}

	/*
		Ejercicio, realizar una funcion en la que podamos identificar si una palabra
		o frase es un palindromo. Ej: Ama o Amor a Roma.
	*/
	isPalidnromo("Amor a Roma")
}
