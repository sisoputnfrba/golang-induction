package main

import "fmt"

func main() {
	funcionCase("A")
	funcionCase("D")
	funcionCase("J")
}

func funcionCase(a string) {
	switch a {
	case "A":
		fmt.Println("Abeja")
	case "B":
		fmt.Println("Baskett")
	case "C":
		fmt.Println("Codigo")
	case "D":
		fmt.Println("Dado")
	default:
		fmt.Println("No se que es!")
	}
}
