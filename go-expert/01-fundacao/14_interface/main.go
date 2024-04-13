package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Ativo      bool
}

type Pessoa interface {
	Desativar()
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

type Empresa struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (e *Empresa) Desativar() {
	e.Ativo = false
	fmt.Printf("A empresa %s foi desativada\n", e.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	fulano := Client{
		Nome:  "Fulano",
		Idade: 30,
		Ativo: true,
	}

	empresa := Empresa{
		Nome:  "Test 123",
		Idade: 10,
		Ativo: true,
	}

	Desativacao(&fulano)
	Desativacao(&empresa)
}
