package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(ValorDoSaque float64) string {
	podeSacar := ValorDoSaque > 0 && ValorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= ValorDoSaque
		return "Saque Realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(ValorDoDeposito float64) (string, float64) {
	if ValorDoDeposito > 0 {
		c.saldo += ValorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito insuficiente", c.saldo
	}

}

func (c *ContaCorrente) Trasnferir(valorDaTrasnferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTrasnferencia < c.saldo && valorDaTrasnferencia > 0 {
		c.saldo -= valorDaTrasnferencia
		contaDestino.Depositar(valorDaTrasnferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
