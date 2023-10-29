package main

import "fmt"

func main() {
	var mapSinValor map[string]int
	// Declarar un map con claves de tipo cadena y valores de tipo entero
	var mapConValor map[string]int = map[string]int{"Juan": 28, "Ana": 30}

	fmt.Println("Map sin valor: ", mapSinValor)
	fmt.Println("Map con valor: ", mapConValor)

	// Verificamos la edad de Juan
	fmt.Println("Valor para Juan: ", mapConValor["Juan"])
	// Verificamos la edad de NoExiste
	fmt.Println("Valor para NoExiste: ", mapConValor["NoExiste"])

	// Agregamos el dato Nahuel y lo imprimimos
	mapConValor["Nahuel"] = 32
	fmt.Println("Valor para Nahuel: ", mapConValor["Nahuel"])

	// Eliminamos a Ana e imprimimos el mapa
	delete(mapConValor, "Ana")
	fmt.Println("Map con valor: ", mapConValor)
}
