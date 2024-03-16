package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Ativo      bool
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	fulano := Client{
		Nome:  "Fulano",
		Idade: 30,
		Ativo: true,
	}

	fulano.Numero = 20
	fulano.Logradouro = "teste 123"
	fulano.Endereco.Ativo = true

	fmt.Println(fulano)
}
