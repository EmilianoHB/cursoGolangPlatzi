package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Emiliano"] = 25
	m["Pepe"] = 35
	m["Claudia"] = 23

	fmt.Println(m)

	//Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor dentro del map
	valor, ok := m["Jose"]
	fmt.Println(valor, ok)
	valor, ok = m["Emiliano"]
	fmt.Println(valor, ok)
}
