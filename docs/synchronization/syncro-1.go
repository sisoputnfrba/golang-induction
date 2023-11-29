package main

import (
	"fmt"
	"sync"
)

var valor int = 0
var wg sync.WaitGroup

func main() {

	for i := 0; i < 20; i++ {
		// Agrego uno al WaitGroup antes de comenzar cada rutina
		wg.Add(1)
		go thread()
	}

	// Esperar a que todas las rutinas finalicen
	wg.Wait()
}

func thread() {
	// Indico que al finalizar la función thread tiene que liberar un valor de wg
	defer wg.Done()

	valor++
	fmt.Println("Este es el hilo número", valor)
}
