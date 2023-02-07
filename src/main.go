package main

import "fmt"

func main() {
	//Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%v tiene mas de %v cursos\n", nombre, cursos)
	fmt.Print(message)

	//Tipos de datos
	fmt.Printf("hellomessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
