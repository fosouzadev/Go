package contas

type ContaInterface interface {
	Sacar(valor float64) bool
}
