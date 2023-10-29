package main

import "fmt"

func main() {
	//Definimos las variables
	var cadena string = "ABC€D"
	var arrayBytes []byte

	// Transformamos nuestra cadena en un array de bytes
	arrayBytes = []byte(cadena)

	// Imprimimos el array
	fmt.Println("Array de bytes:", arrayBytes)

	// Convertimos el array de bytes en una cadena
	var nuevaCadena string = string(arrayBytes)

	// Imprimimos la nueva cadena
	fmt.Println("Nueva cadena:", nuevaCadena)
}
