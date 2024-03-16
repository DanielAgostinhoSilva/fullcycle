package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(10, 20, 30, 40, 50))
}

func sum(valores ...int) int {

	total := 0
	for _, numero := range valores {
		total += numero
	}
	return total
}
