package mypackage

import "fmt"

// Pc es un strunc de una pc
type Pc struct {
	Ram   int
	Disco int
	Marca string
}

func (miPc Pc) Ping() {
	fmt.Println(miPc.Marca, "Pong")
}

// Funcion para modificar atributos utilizando el puntero
func (miPc *Pc) DuplicarRam() {
	miPc.Ram += miPc.Ram
}

// Personalizacion de textos con Sprintf
func (miPc Pc) String() string {
	return fmt.Sprintf("Tengo una computadora de marca %s, que tiene %d de memoria RAM y un disco de %d GB", miPc.Marca, miPc.Ram, miPc.Disco)
}
