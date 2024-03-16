package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

	valor, err = sum(10, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, error) {

	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
