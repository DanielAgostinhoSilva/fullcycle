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
	fmt.Printf("O tipo de A é %T seu valor %v\n", a, a)
	fmt.Printf("O tipo de B é %T seu valor %v\n", b, b)
	fmt.Printf("O tipo de C é %T seu valor %v\n", c, c)
	fmt.Printf("O tipo de D é %T seu valor %v\n", d, d)
	fmt.Printf("O tipo de E é %T seu valor %v\n", e, e)
	fmt.Printf("O tipo de F é %T seu valor %v\n", f, f)
}
