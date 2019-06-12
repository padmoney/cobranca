package cobranca

type Conta struct {
	Banco         string
	Agencia       string
	ContaCorrente string
	Carteira      string
	Variacao      string
	Convenio      string
}

func NewConta(banco string, agencia, contaCorrente, carteira, convenio string) Conta {
	return Conta{
		Banco:         banco,
		Agencia:       agencia,
		ContaCorrente: contaCorrente,
		Carteira:      carteira,
		Convenio:      convenio}
}
