package main

import "fmt"

func main() {
	var valor1 int = 20
	var valor2 int = 8

	imprimir(valor1)
	imprimir(valor2)
}

func imprimir(a int) {
	fmt.Println("El valor es:", a)
}
