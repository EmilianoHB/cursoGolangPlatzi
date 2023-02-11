package mypackage

/*Creamos una funcion que agarre un mensaje y lo guarde en un canal*/
func Mensaje(text string, c chan<- string) {
	c <- text
}
