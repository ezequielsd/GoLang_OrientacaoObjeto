package contas

import "bancoEzequiel/clientes"

type ContaPoupanca struct{
	Titular clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return  "Saque realizado com sucesso R$", c.saldo
	}else{
		return "saldo insuficiente R$", c.saldo
	}	
}

func (c *ContaPoupanca) Depositar (valorDoDeposito float64) (string, float64){
	
	if valorDoDeposito > 0{
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso. saldo: R$", c.saldo
	} else {
		return "Depósito não pode ser realizado. saldo: R$", c.saldo
	}	
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
