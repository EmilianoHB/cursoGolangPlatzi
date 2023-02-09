package main

import (
	pk "cursoGolangPlatzi/src/mypackage"
	"fmt"
)

func main() {
	cuadradito := pk.Cuadrado{
		Base: 3,
	}
	rectangulito := pk.Rectangulo{
		Base:   2,
		Altura: 5,
	}

	pk.Calcular(cuadradito)
	pk.Calcular(rectangulito)

	//Lista de interfaces
	miInterfaz := []interface{}{
		"Hola",
		12,
		4.90,
		true,
	}
	fmt.Println(miInterfaz...)
}
