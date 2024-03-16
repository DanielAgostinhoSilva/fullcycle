package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(10, 20, 30, 40, 50) * 2
	}()

	fmt.Println(total)
}

func sum(valores ...int) int {

	total := 0
	for _, numero := range valores {
		total += numero
	}
	return total
}
