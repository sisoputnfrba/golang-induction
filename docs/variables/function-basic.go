package main

import "fmt"

func main() {
	var valor1 int = 1
	var valor2 int = 2

	var suma int = sumar(valor1, valor2)

	fmt.Println("Resultado:", suma)
}

func sumar(a int, b int) int {
	return a + b
}
