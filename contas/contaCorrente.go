package contas

import (
	"bancoEzequiel/clientes"	
)

type ContaCorrente struct{
	Titular clientes.Titular 
	NumeroAgencia, NumeroConta int 
	saldo float64 
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return  "Saque realizado com sucesso R$" 
	}else{
		return "saldo insuficiente R$"
	}	
}

func (c *ContaCorrente) Depositar (valorDoDeposito float64) (string, float64){
	
	if valorDoDeposito > 0{
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso. saldo: R$", c.saldo
	} else {
		return "Depósito não pode ser realizado. saldo: R$", c.saldo
	}	
}

func (c *ContaCorrente) Transferencia(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}