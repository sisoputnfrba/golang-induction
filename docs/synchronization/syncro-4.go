package main

import (
	"fmt"
	"sync"
)

var valores []int
var valor = 0
var wg sync.WaitGroup
var mutex sync.Mutex

// Incializamos el semaforo del productor en 3 para tener hasta 3 recursos en el mercado
var sem_productor = make(chan int, 3)

// Inicializamos el semaforo consumidor en 0 ya que no tenemos recursos en el mercado
var sem_consumidor = make(chan int)

func main() {

	// Ejecutamos 5 productor
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go productor()
	}

	// Ejecutamos 5 consumidores
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go consumidor()
	}

	// Esperamos a que finalicen todas las rutinas
	wg.Wait()
}

func productor() {
	defer wg.Done()

	for true {
		// Consumimos un semaforo de productor que permite hasta 3 recursos
		sem_productor <- 0

		// Entramos en la sección critica de nuestro mercado, Las variables valor y valores
		mutex.Lock()
		valores = append(valores, valor)
		fmt.Println("Producimos el valor ", valor)
		valor++
		mutex.Unlock()

		// Informamos a consumidores que tienen un recurso en el mercado
		<-sem_consumidor
	}
}

func consumidor() {
	defer wg.Done()

	for true {
		// Tomo un recurso del mercado
		sem_consumidor <- 0

		// Entro en mi sección critica de la variable valores
		mutex.Lock()
		fmt.Println("Consumimos el valor ", valores[0])
		valores = valores[1:]
		mutex.Unlock()

		// Aviso al productor q tiene un lugar libre en el mercado
		<-sem_productor
	}
}
