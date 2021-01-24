package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// cria como estrutura
	contaGuilherme := ContaCorrente{titular: "Guilherme", saldo: 10000}

	contaBruna := ContaCorrente{"Bruna", 222, 111222, 2000}

	// cria como um objeto que pode ser nil, é um ponteiro
	contaCris := new(ContaCorrente)
	contaCris.titular = "Cris"

	var contaPaulo *ContaCorrente
	contaPaulo = new(ContaCorrente)
	contaPaulo.titular = "Paulo"

	//contaPaulo = nil  // permite
	//contaBruna = nil  // não permite

	fmt.Println(contaGuilherme)
	fmt.Println(contaBruna)
	fmt.Println("Ponteiro", contaCris, "Valor do ponteiro", *contaCris)
	fmt.Println("Ponteiro", contaPaulo, "Valor do ponteiro", *contaPaulo)
}
