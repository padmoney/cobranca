package boleto

import "github.com/padmoney/cobranca"

func NewConta(banco string, agencia, contaCorrente, carteira, convenio string) cobranca.Conta {
	return cobranca.Conta{
		Banco:         banco,
		Agencia:       agencia,
		ContaCorrente: contaCorrente,
		Carteira:      carteira,
		Convenio:      convenio}
}
