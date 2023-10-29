package main

import "fmt"

var valor1 int = 3

func main() {
	var valor1 int = 1
	var valor2 int = 2

	fmt.Println("Resultado sumar:", sumar(valor1, valor2))
	fmt.Println("Resultado sumar2:", sumar2(valor2))
}

func sumar(a int, b int) int {
	return a + b
}

func sumar2(b int) int {
	return valor1 + b
}
