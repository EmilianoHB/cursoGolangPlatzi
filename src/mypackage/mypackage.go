package mypackage

/*
Creamos una funcion llamada Say() donde va a recibir un texto en forma
de string, agregandole el canal. Ej:

	func Say(text string, c chan string) {
		c <- text
	}

Una buena practica es indicar si el parametro del canal es de entrada o de salida,
mejorando tu practica y la eficiencia del codigo. Ej:

 1. Este canal que vamos a agregar solo va a ser de entrada de datos

	func Say(text string, c chan<- string) {
	    c <- text
	}

 2. Este canal que vamos a agregar solo va a ser de salida de datos, guardando
    en la variable text1

	func Say(text string, c <-chan string) {
	    text1 <- c
	}
*/
func Say(text string, c chan<- string) {
	/*Tomamos el texto y lo pasamos al canal
	c <- simbolo de igresar datos al canal*/
	c <- text
}
