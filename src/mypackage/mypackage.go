package mypackage

import (
	"fmt"
	"sync"
)

func Say(text string, wg *sync.WaitGroup) {
	defer wg.Done() /* Esta linea se va a ejecutar hasta el final
	de la funcion, y de esta forma libera el gorutine del WaitGroup*/
	fmt.Println(text)
}
