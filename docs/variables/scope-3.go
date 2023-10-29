package main

import "fmt"

func main() {
	var valor1 int = 1
	var valor2 int = 2

	var sumar = func(a int, b int) int {
		return a + a + b
	}

	fmt.Println("Resultado sumar:", sumar(valor1, valor2))
}

func sumar(a int, b int) int {
	return a + b
}
