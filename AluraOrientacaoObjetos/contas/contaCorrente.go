package contas

import "Go/AluraOrientacaoObjetos/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) bool {
	if valor >= 0 && valor <= c.saldo {
		c.saldo -= valor
		return true
	}

	return false
}

func (c *ContaCorrente) Depositar(valor float64) (bool, float64) {
	if valor > 0 {
		c.saldo += valor
		return true, c.saldo
	}

	return false, c.saldo
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	}

	return false
}

func (c ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
