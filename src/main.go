package main

import (
	pk "cursoGolangPlatzi/src/mypackage"
	"fmt"
)

func main() {
	/*Creamos nuestro primer canal, vamos a manejar dos datos a la vez*/
	c := make(chan string, 2)
	/*Insertamos dos mensajes*/
	c <- "mensaje1"
	c <- "mensaje2"

	/*Nos introducirmos en conceptos utiles al manejar los channels
	len()(Esta funcion nos dice la cantidad de datos tenemos dentro del channel)
	y cap()( Nos indica la cantidad maxima de datos que puede almacenar
	ese channel)*/
	fmt.Println(len(c), cap(c))

	/*Close
	La funcion close() indica al runtime de go que va a cerrar el canal, indicando
	que ese canal no va a recibir ningun dato adicional, a pesar de que tenga o no
	mas capacidad. Lo ideal es cerrar los canales una vez sabes que no va a
	recibir mas datos, no solo es una buena practica sino que tambien
	mejora la eficiencia del codigo a nivel del runtime de go*/
	close(c)

	/*Range
	Nos ayuda a hacer un recorrido de todos los mensajes que estan en este channel
	por ejemplo. Es ideal cuando queremos iterar en cada uno de los elementos de
	un canal que usualmente puede estar abierto o incluso manejar datos que en
	algun punto pueden ser desconocidos*/
	for mensaje := range c {
		fmt.Println(mensaje)
	}

	/*Cuando empezamos a manejar multiples canales y no tenemos la certeza
	de cual va a responder primero es cuando empezamos a utilizar select.
	Para ejemplificar vamos a crear dos canales que van a recibir emails

	Select

	No le indicaremos la capacidad maxima, sino que los dejaremos de
	forma dinamica*/
	email1 := make((chan string))
	email2 := make((chan string))

	/*Invocaremos la funcion Mensaje con una goroutine de manera independiente.*/
	go pk.Mensaje("Mensaje 1", email1)
	go pk.Mensaje("Mensaje 2", email2)

	/*En eset punto no tenemos a ciencia cierta cual de los dos canales va a
	responder primero y ahi es donde empezamos a utlilizar select. Para eso
	es importante que tengas precente la cantidad de datos que vas a manejar
	en cada uno de los channels, icluso la cantidad de channels que vas a
	tener por que los vamos a manejar a travez de un for*/
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
