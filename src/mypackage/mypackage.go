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
func (miPc *Pc) DuplicarRam() {
	miPc.Ram += miPc.Ram
}
