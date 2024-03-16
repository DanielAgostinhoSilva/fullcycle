package main

import "fmt"

func main() {
	salarios := map[string]int{"Fulano 1": 1000, "Fulano 2": 2000, "Fulano 3": 2000}

	showMap(salarios)

	delete(salarios, "Fulano 1")

	showMap(salarios)

	salarios = make(map[string]int)
	salarios["Fulano 4"] = 3000
	showMap(salarios)
}

func showMap(salarios map[string]int) {
	for i, v := range salarios {
		fmt.Printf("O salário de %s é R$ %d\n", i, v)
	}
}
