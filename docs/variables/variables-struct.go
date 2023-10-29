package main

import "fmt"

// Declarar una estructura con diferentes campos
type Persona struct {
	nombre string
	edad   int
	altura float32
}

func main() {
	var structSinValor Persona
	var structConValor Persona = Persona{"Nahuel", 32, 1.77}

	fmt.Println("Struct sin valor:", structSinValor)
	fmt.Println("Struct con valor:", structConValor)

	// Para acceder a una única propiedad
	fmt.Println("Nombre del struct:", structConValor.nombre)

	// Para redefinir una propiedad
	structConValor.edad = 33
	fmt.Println("Edad del struct:", structConValor.edad)
}
