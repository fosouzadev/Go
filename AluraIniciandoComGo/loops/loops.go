package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3}
	indice := 0

	// loop infinito
	for {
		fmt.Println("indice", indice, "valor", numeros[indice])
		indice++

		if indice == len(numeros) {
			break
		}
	}

	fmt.Println()

	// loop tradicional
	//for i := 0; i < len(numeros); i++ {
	for indice = 0; indice < len(numeros); indice++ {
		fmt.Println("indice", indice, "valor", numeros[indice])
	}

	fmt.Println()

	// for range
	// for i, value := range numeros {
	var valor int
	for indice, valor = range numeros {
		fmt.Println("indice", indice, "valor", valor)
	}

}
