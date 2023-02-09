package mypackage

import "fmt"

/*
Los structs o funciones se diferencian entre publicas y privadas dependiendo de las
iniciales estan como minusculas o mayusculas
Si la letra inicial es mayuscula es  publica.
Si la letra inicial esta en minuscula es privada.
*/

// AutoPublico Clase Auto con acceso publico
type AutoPublico struct {
	Marca    string
	Modelo   int
	Color    string
	Puertas  int
	Asientos int
}

func ImprimirMensaje(text string) {
	fmt.Println(text)
}
