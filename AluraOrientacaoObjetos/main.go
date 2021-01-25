package main

import (
	cli "Go/AluraOrientacaoObjetos/clientes"
	"Go/AluraOrientacaoObjetos/contas"
	"fmt"
)

func main() {
	// cria como estrutura
	contaGui := contas.ContaCorrente{Titular: cli.Titular{Nome: "Guilherme"}, NumeroAgencia: 12, NumeroConta: 123}
	contaGui.Depositar(100)

	contaCris := contas.ContaPoupanca{Titular: cli.Titular{Nome: "Cris"}, NumeroAgencia: 14, NumeroConta: 525}
	contaCris.Depositar(300)

	if contaCris.Sacar(40) {
		fmt.Println("Saque ok, saldo : ", contaCris.ObterSaldo())
	} else {
		fmt.Println("Valor inválido, saldo : ", contaCris.ObterSaldo())
	}

	_, saldoAtual := contaCris.Depositar(40)
	fmt.Println(saldoAtual)

	if contaCris.Transferir(50, &contaGui) {
		fmt.Println("Transferencia ok", contaGui, contaCris)
	} else {
		fmt.Println("Transferencia inválida", contaGui, contaCris)
	}
}
