package main

import "fmt"

func main() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var option int
	//fmt.Scanf("%d", &option) // ou
	fmt.Scan(&option)

	if option == 1 {
		fmt.Println("Iniciando monitoramento")
	} else if option == 2 {
		fmt.Println("Exibindo logs")
	} else if option == 3 {
		fmt.Println("Saindo")
	} else {
		fmt.Println("Opção inválida")
	}
}
