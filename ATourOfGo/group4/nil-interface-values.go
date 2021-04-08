package main

import "fmt"

type I interface {
	M()
}

func main() {
	// variavel de interface que recebe seja nil gera erro ao chamar metodos
	// pois não possui um tipo concreto associado
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
