package main

import (
	"fmt"
	"bancoEzequiel/contas"
	"bancoEzequiel/clientes"
)


func PagarBoleto(conta verificarConta, valorDoBoleto float64){
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface{
	Sacar(valor float64) string
}

func main() {
	
	fmt.Println("Conta Corrente ========================================== ")

	contaDoEzequiel := contas.ContaCorrente{Titular: clientes.Titular{
		Nome: "Ezequiel",
		CPF: "011.111.111-11",
		Profissao: "Engenheiro",
	},
	NumeroAgencia: 911,
	NumeroConta: 56984 }

	contaDoEzequiel.Depositar(100)
	fmt.Println("Depositado valor de R$", contaDoEzequiel.ObterSaldo())
	valorBoleto := 65.00
	PagarBoleto(&contaDoEzequiel, valorBoleto)
	fmt.Println("Boleto pago no valor de R$", valorBoleto, " novo saldo de ", contaDoEzequiel.ObterSaldo())

	fmt.Println(contaDoEzequiel)
	fmt.Println("Saldo da conta de", contaDoEzequiel.Titular.Nome, "é de ", contaDoEzequiel.ObterSaldo())

	fmt.Println("Conta Poupanca ========================================== ")

	contaDaLambreta := contas.ContaPoupanca{Titular: clientes.Titular{
		Nome: "Lambreta",
		CPF: "022.222.222-22",
		Profissao: "Assassino de aluguel",
	},
	NumeroAgencia: 912,
	NumeroConta: 40521,
	Operacao: 13 }

	contaDaLambreta.Depositar(200)

	fmt.Println(contaDaLambreta)
	fmt.Println("Saldo da conta de", contaDaLambreta.Titular.Nome, "é de ", contaDaLambreta.ObterSaldo())
}

