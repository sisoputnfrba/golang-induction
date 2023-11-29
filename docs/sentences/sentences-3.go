package main

import "fmt"

func main() {
	var edad int = 76
	if edad < 18 {
		fmt.Println("Menor de edad!")
	} else if edad < 75 {
		fmt.Println("Mayor de edad!")
	} else {
		fmt.Println("Jubilado")
	}
}
