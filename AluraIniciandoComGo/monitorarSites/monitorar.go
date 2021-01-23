package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 5

func main() {
	for { // significa um loop infinito, podendo finalizar com break
		exibirMenu()
		option := lerOpcao()
		executar(option)
	}
}

func exibirMenu() {
	fmt.Println()
	fmt.Println("------ Menu ------")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Iniciar monitoramento automático")
	fmt.Println("0 - Sair do programa")
}

func lerOpcao() int {
	var option int
	fmt.Scan(&option) //fmt.Scanf("%d", &option) // ou
	fmt.Println()

	return option
}

func executar(option int) {
	switch option {
	case 1:
		fmt.Println("Iniciando monitoramento")
		iniciarMonitoramento(false)
	case 2:
		fmt.Println("Exibindo logs")
		exibindoLogs()
	case 3:
		fmt.Println("Iniciando monitoramento")
		iniciarMonitoramento(true)
	case 0:
		fmt.Println("Saindo")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida")
		os.Exit(-1)
	}

	/*if option == 1 {
		fmt.Println("Iniciando monitoramento")
	} else if option == 2 {
		fmt.Println("Exibindo logs")
	} else if option == 0 {
		fmt.Println("Saindo")
	} else {
		fmt.Println("Opção inválida")
	} */
}

func iniciarMonitoramento(automatico bool) {
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://alura.com.br",
		"https://caelum.com.br",
	}
	//sites := []string{"https://random-status-code.herokuapp.com/", "https://alura.com.br", "https://caelum.com.br"}

	if automatico {
		for {
			for _, site := range sites {
				monitorarSite(site)
			}

			time.Sleep(delay * time.Second)
			fmt.Println()
		}
	} else {
		for _, site := range sites {
			monitorarSite(site)
		}
	}
}

// func leSitesDoArquivo() []string {
// 	arquivo, _ := os.Open("sites.txt")
// 	arquivo.
// }

func monitorarSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code: ", response.StatusCode)
	}
}

func exibindoLogs() {

}
