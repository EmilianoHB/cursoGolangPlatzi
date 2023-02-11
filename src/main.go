package main

import (
	pk "cursoGolangPlatzi/src/mypackage"
	"fmt"
)

func main() {
	/*Para crear un channel, creamos el nombre de una variable. Utilizaremos
	la funcion make() (que solamente se utiliza para slice, map y channels),
	indicando que vamos a crear un chan, al escribirlo tenemos que indicar
	el tipo de dato que va a pasar por ese canal. El siguiente paso es opcional,
	pero a nivel de optimizacion es una buena practica, en el cual despues
	de una coma, vamos a indicarle cuantos datos simultaneos va a manejar ese
	canal, sin embargo, si no agregamos ese dato, el canal va a tomar un valor
	dinamico en todo momento*/
	c := make(chan string, 1)

	//Para funciones practicas creamos un hello
	fmt.Println("Hello")

	/*Invocamos la funcion con una goroutin. Especificamos el texto y le
	agregamos el canal*/
	go pk.Say("Bye", c)

	/*Para asegurarnos que la goroutin del main espere a la ejecucion de
	esta funcion antes de terminar, lo hacemos extrayendo el dato que agregamos
	en el canal anteriormente con un PL
		<-c simbolo de salida del dato del canal*/
	fmt.Println(<-c)
}
