package main

import "fmt"

func main() {
	ola()
	fmt.Println(olaAlguem())
	olaPessoa("Oi", "Fulano")
	fmt.Println(retornaOlaPessoa("Oi", "Fulano"))

	nome, idade := maisDeUmRetorno()
	_, idade2 := maisDeUmRetorno() // podemos ignorar um dos retornos com _ underline
	fmt.Println(nome, idade, idade2)
}

// sem parametro / sem retorno
func ola() {
	fmt.Println("Olá")
}

// sem paramero / com retorno
func olaAlguem() string {
	return "Olá alguem"
}

// com parametro / sem retorno
func olaPessoa(saudacao string, nome string) {
	fmt.Println(saudacao, nome)
}

// com parametro / com retorno
func retornaOlaPessoa(saudacao string, nome string) string {
	frase := saudacao + " " + nome
	return frase
}

// mais de um retorno
func maisDeUmRetorno() (string, int) {
	nome := "Fulano"
	idade := 30
	return nome, idade
}
