package main

import "fmt"

func main() {
	array()
	slices()
}

// tem tamanho fixo na sua declaracao
func array() {
	fmt.Println()
	fmt.Println("Array")

	var nomes1 [3]string = [3]string{"joao", "maria", "paulo"}

	var nomes2 [3]string
	nomes2[0] = "joao"
	nomes2[1] = "maria"
	nomes2[2] = "paulo"

	nomes3 := [3]string{"joao", "maria", "paulo"}
	// nomes3 = append(nomes3, "cesar") não permite porque não é um slice

	var nomes4 []string = nomes3[1:2] // consegue extrair dados de um array para um slice

	fmt.Println(nomes1, "tamanho", len(nomes1), "capacidade", cap(nomes1))
	fmt.Println(nomes2, "tamanho", len(nomes2), "capacidade", cap(nomes2))
	fmt.Println(nomes3, "tamanho", len(nomes3), "capacidade", cap(nomes3))
	fmt.Println(nomes4, "tamanho", len(nomes4), "capacidade", cap(nomes4))
}

// é uma abstração do array com tamanho definido em sua atribuição de valores
// permite adicionar novos elementos, e quando a quantidade ultrapassa a capacidade atual
// sua capacidade é dobrada
func slices() {
	fmt.Println()
	fmt.Println("Slice")

	var nomes1 []string = []string{"joao", "maria", "paulo"}

	var nomes2 []string
	nomes2 = []string{"joao", "maria", "paulo"}

	nomes3 := []string{"joao", "maria", "paulo"}
	nomes3 = append(nomes3, "cesar") // dobra capacidade de 3 para 6, e fica com tamanho 4 por ter 4 nomes

	var nomes4 []string = nomes3[1:4]

	fmt.Println(nomes1, "tamanho", len(nomes1), "capacidade", cap(nomes1))
	fmt.Println(nomes2, "tamanho", len(nomes2), "capacidade", cap(nomes2))
	fmt.Println(nomes3, "tamanho", len(nomes3), "capacidade", cap(nomes3))
	fmt.Println(nomes4, "tamanho", len(nomes4), "capacidade", cap(nomes4))
}
