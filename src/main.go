package main

import "fmt"

func normalFuncion(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFuncion("HOLA MUNDO")
	tripleArgument(1, 2, "Emiliano")

	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Valores: ", value1, value2)

	//Para utilizar solamente uno de los dos valores retomados
	value1, _ = doubleReturn(2)
	fmt.Println("Valor: ", value1)
}
