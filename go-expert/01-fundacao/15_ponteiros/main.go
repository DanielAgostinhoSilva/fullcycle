package main

func main() {
	// momoria -> endereco -> valor

	a := 10
	println(&a)
	var ponteiro *int = &a
	println(ponteiro)
	*ponteiro = 20
	println(a)
	b := &a
	*b = 10
	println(a)
	println(*b)
}
