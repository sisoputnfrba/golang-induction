package main

import (
	"fmt"
	"sync"
)

var valor int = 0
var wg sync.WaitGroup
var mutex sync.Mutex

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
	mutex.Lock()
	valor++
	fmt.Println("Este es el hilo número", valor)
	//Soltamos el candado de mutex
	mutex.Unlock()
}
