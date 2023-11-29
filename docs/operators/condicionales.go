package main

import "fmt"

func main() {
	var azul bool = true
	var amarillo bool = true
	var rojo bool = false

	fmt.Println("Verde = ", azul && amarillo)
	fmt.Println("Violeta = ", azul && rojo)
	fmt.Println("Naranja = ", amarillo && rojo)

	fmt.Println("Daltonico = ", amarillo || rojo || azul)
	fmt.Println("No me juzguen, soy Daltonico")
}
