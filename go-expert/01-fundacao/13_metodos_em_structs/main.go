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

func (c *Client) Desativar() {
	c.Ativo = false
}

func main() {
	fulano := Client{
		Nome:  "Fulano",
		Idade: 30,
		Ativo: true,
	}

	fulano.Desativar()

	fulano.Numero = 20
	fulano.Logradouro = "teste 123"
	fulano.Endereco.Ativo = true

	fmt.Println(fulano)
}
