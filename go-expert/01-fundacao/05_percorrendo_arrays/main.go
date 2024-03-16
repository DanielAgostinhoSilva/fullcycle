package main

import "fmt"

const (
	a = "Hello World"
)

type ID int

var (
	b bool = true
	c int
	d string = "teste d"
	e float64
	f ID = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(len(meuArray))
	fmt.Println("O indice do ultimo elemento é", len(meuArray)-1)
	fmt.Println("O valor do ultimo elemento é", meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}

}
