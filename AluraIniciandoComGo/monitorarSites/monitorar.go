package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	// sites := []string{
	// 	"https://random-status-code.herokuapp.com/",
	// 	"https://alura.com.br",
	// 	"https://caelum.com.br",
	// }
	//sites := []string{"https://random-status-code.herokuapp.com/", "https://alura.com.br", "https://caelum.com.br"}
	sites := leSitesDoArquivo()

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

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, erro := os.Open("sites.txt") // os package
	//arquivo, erro := ioutil.ReadFile("sites.txt") // io/ioutil package

	if erro != nil {
		fmt.Println("Erro ao ler sites: ", erro)
		return sites
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if erro == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func monitorarSite(site string) {
	response, erro := http.Get(site)

	if erro != nil {
		fmt.Println("Erro no Get ao site: ", erro)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code: ", response.StatusCode)
		registraLog(site, false)
	}
}

func registraLog(site string, online bool) {
	arquivo, erro := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if erro != nil {
		fmt.Println("Erro ao registrar log: ", erro)
		return
	}

	// observar formato de data/hora no site https://golang.org/src/time/format.go
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(online) + "\n")

	arquivo.Close()
}

func exibindoLogs() {
	// não precia fechar o arquivo
	// auxilia a imprimir arquivo inteiro, sem precisar ler linha a linha
	arquivo, erro := ioutil.ReadFile("log.txt")

	if erro != nil {
		fmt.Println("Erro ao ler logs: ", erro)
		return
	}

	fmt.Println(string(arquivo))
}
