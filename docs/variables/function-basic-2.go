package main

import "fmt"

func main() {
	var valor1 int = 20
	var valor2 int = 8

	fmt.Println("Dividir Enteros 20/8:", divideInts(valor1, valor2))

	fmt.Println("Dividir 20/8:", divideFloat(valor1, valor2))
	fmt.Println("Dividir 8/20:", divideFloat(valor2, valor1))
}

func divideInts(a int, b int) int {
	return a / b
}

func divideFloat(a int, b int) float32 {
	return float32(a) / float32(b)
}
