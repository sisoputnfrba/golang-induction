package main

import (
	"fmt"
	"sync"
)

var valor int = 0
var wg sync.WaitGroup
var sem = make(chan int, 1)

func main() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go thread()
	}

	wg.Wait()
}

func thread() {
	defer wg.Done()

	//Obtenemos el candado de mutex
	sem <- 0
	valor++
	fmt.Println("Este es el hilo número", valor)
	//Soltamos el candado de mutex
	<-sem
}
