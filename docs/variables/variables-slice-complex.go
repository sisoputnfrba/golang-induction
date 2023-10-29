package main

import "fmt"

func main() {
	var mySlice []int = []int{1, 2, 3, 4, 5}

	// Para obtener los primeros dos elementos del array
	fmt.Println("Primeros dos elementos:", mySlice[:2])

	// Para quedarme con todos los elementos posterior al indice 2 (inclusive)
	fmt.Println("Ultimos elementos:", mySlice[2:])

	// Para quitar el elemento de indice 2
	fmt.Println("Slice sin el elemento 3:", append(mySlice[:2], mySlice[2+1:]...))
}
