package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	fulano := Client{
		Nome:  "Fulano",
		Idade: 30,
		Ativo: true,
	}

	fmt.Println(fulano)
}
