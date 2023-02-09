package mypackage

import "fmt"

// Creamos un struct de cuadrado
type Cuadrado struct {
	Base float64
}

// Creamos una funcion del caluclo del area del cuadrado
func (c Cuadrado) Area() float64 {
	return c.Base * c.Base
}

// Creamos un struct de rectangulo
type Rectangulo struct {
	Base   float64
	Altura float64
}

// Creamos la funcion del calculo del cuadrado
func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

/*
Creamos un interface
Lo que hace unicamente es tomar la funcion que comparten los dos
Struct
*/
type Figuras2D interface {
	Area() float64
}

/*
Funcion calcular lo que hace es tomar como entrada la interfaz
que creamos antes y hacemos que ejecute la funcion de area
*/
func Calcular(f Figuras2D) {
	fmt.Println("Area: ", f.Area())
}
