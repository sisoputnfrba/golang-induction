package main

import "fmt"

func main() {
	var i int = 10
	i += 5

	fmt.Println("i += 5 is ", i)

	i -= 3
	fmt.Println("i -= 3 is ", i)

	i *= 4
	fmt.Println("i *= 4 is ", i)

	i /= 2
	fmt.Println("i /= 2 is ", i)
}
