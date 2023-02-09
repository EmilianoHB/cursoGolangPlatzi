package main

import (
	"cursoGolangPlatzi/src/mypackage"
	"fmt"
)

func main() {
	a := 50
	b := &a
	fmt.Println(a) //variable
	fmt.Println(b) //puntero de la variable
	fmt.Println(b)
	fmt.Println(*b) //saber a que le apunta el puntero
	*b = 100        //si modificamos el puntero
	fmt.Println(a)  //modificaremos el valor de la variable

	miPc := mypackage.Pc{
		Ram:   16,
		Disco: 200,
		Marca: "Msi",
	}
	fmt.Println(miPc)

	miPc.Ping()

	//Una forma de utilizar la ram usando punteros
	fmt.Println(miPc)
	miPc.DuplicarRam()

	fmt.Println(miPc)
	miPc.DuplicarRam()

	fmt.Println(miPc)
}
