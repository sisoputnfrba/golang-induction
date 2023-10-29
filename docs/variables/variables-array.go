package main

import "fmt"

func main() {
	var arraySinValor [5]int
	var arrayConValor [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array sin valor: ", arraySinValor)
	fmt.Println("Array con valor: ", arrayConValor)

	// El indice comienza desde la posicion 0, por lo que el valor 3 es la posición 2 del array
	fmt.Println("Valor 3 del array: ", arrayConValor[2])
}
