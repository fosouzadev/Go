package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println("Formas de criar um objeto struct")
	pessoa1 := Pessoa{}
	pessoa2 := Pessoa{nome: "Juca"}
	pessoa3 := Pessoa{"Cesar", 40}
	var pessoa4 = Pessoa{idade: 25}
	var pessoa5 Pessoa = Pessoa{nome: "July", idade: 20}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)
	fmt.Println(pessoa4)
	fmt.Println(pessoa5)

	fmt.Println("\nFormas de criar um objeto ponteiro struct")
	pessoa6 := new(Pessoa)
	pessoa6.nome = "Jao"

	var pessoa7 *Pessoa
	pessoa7 = new(Pessoa)
	pessoa7.nome = "Noe"

	fmt.Println(pessoa6)
	fmt.Println(pessoa7)

	fmt.Println("\nComparacao de struct é por valor das propriedades")
	pessoa8 := Pessoa{nome: "Sem", idade: 400}
	pessoa9 := Pessoa{nome: "Sem", idade: 400}
	pessoa10 := Pessoa{nome: "Sem", idade: 399}

	fmt.Println("pessoa8 = pessoa9", pessoa8 == pessoa9)
	fmt.Println("pessoa8 = pessoa10", pessoa8 == pessoa10)

	fmt.Println("\nComparacao de ponteiro struct é por valor tambem, porém necessário *")
	pessoa11 := new(Pessoa)
	pessoa11.nome = "Sem"
	pessoa11.idade = 400

	pessoa12 := new(Pessoa)
	pessoa12.nome = "Sem"
	pessoa12.idade = 400

	pessoa13 := new(Pessoa)
	pessoa13.nome = "Sem"
	pessoa13.idade = 400

	fmt.Println("pessoa11 = pessoa12", pessoa11 == pessoa12)
	fmt.Println("pessoa11 = pessoa13", *pessoa11 == *pessoa13)

	fmt.Println("\nComparacao de ponteiro struct com objeto struct, necessário * no ponteiro")

	fmt.Println("pessoa8 == pessoa11", pessoa8 == *pessoa11)
}
