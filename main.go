package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	fmt.Println(contaDoDenis.ObterSaldo())

	PagarBoleto(&contaDoDenis, 50)
	fmt.Println(contaDoDenis.ObterSaldo())
}
