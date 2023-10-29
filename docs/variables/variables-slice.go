package main

import "fmt"

func main() {
	var sliceSinValor []int
	var sliceConValor []int = []int{1, 2, 3, 4, 5}

	fmt.Println("Slice sin valor: ", sliceSinValor)
	fmt.Println("Slice con valor: ", sliceConValor)

	// Imprimo el tamaño del slice con valor
	var tamañoSliceConValor = len(sliceConValor)
	fmt.Println("Tamaño del slice con valor: ", tamañoSliceConValor)

	// Agrego un elemento al slice
	sliceConValor = append(sliceConValor, 6)

	// Imprimo el nuevo tamaño del slice con valor
	tamañoSliceConValor = len(sliceConValor)
	fmt.Println("Nuevo tamaño del slice con valor: ", tamañoSliceConValor)

	// Imprimo el nuevo Slice
	fmt.Println("Nuevo slice con valor: ", sliceConValor)

	// El indice comienza desde la posicion 0, por lo que el valor 3 es la posición 2 del slice
	fmt.Println("Valor 3 del slice: ", sliceConValor[2])
}
