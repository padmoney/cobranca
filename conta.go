package cobranca

import "strings"

type Conta struct {
	Banco         string
	Agencia       string
	ContaCorrente string
	Carteira      string
	Variacao      string
	Convenio      string
	EspecieTitulo string
	Aceite        bool
	Beneficiario  Beneficiario
}

func NewConta(banco string, agencia, contaCorrente, carteira, convenio string, beneficiario Beneficiario) Conta {
	return Conta{
		Banco:         banco,
		Agencia:       agencia,
		ContaCorrente: contaCorrente,
		Carteira:      carteira,
		Convenio:      convenio,
		Beneficiario:  beneficiario}
}

// Aceite do título
// N - Sem aceite
// A - Com aceite - Indica o reconhecimento formal (assinatura no documento) do sacado no título
func (c Conta) GetAceite() string {
	if c.Aceite {
		return "A"
	}
	return "N"
}

// GetEspecieTitulo retorna a espécie de título de acordo com cada banco
// Banco do Brasil
// 01 - Duplicata Mercantil
// 02 - Nota Promissória
// 03 - Nota de Seguro
// 05 - Recibo
// 08 - Letra de Câmbio
// 09 - Warrant
// 10 - Cheque
// 12 - Duplicata de Serviço
// 13 - Nota de Débito
// 15 - Apólice de Seguro
// 25 - Dívida Ativa da União
// 26 - Dívida Ativa de Estado
// 27 - Dívida Ativa de Município
//
// Santander - 033
// 01 - DUPLICATA
// 02 - NOTA PROMISSÓRIA
// 03 - APÓLICE / NOTA DE SEGURO
// 05 - RECIBO
// 06 - DUPLICATA DE SERVIÇO
// 07 - LETRA DE CAMBIO
// 08 - BDP - BOLETO DE PROPOSTA - ( NOTA 6)
// 19 - BCC – BOLETO CARTÃO DE CRÉDITO ( NOTA 8)
func (c Conta) GetEspecieTitulo() string {
	if c.EspecieTitulo == "" {
		return "01"
	}
	return c.EspecieTitulo
}

func (c Conta) NumeroAgencia() string {
	a := strings.Split(c.Agencia, "-")
	if len(a) > 0 {
		return a[0]
	}
	return ""
}

func (c Conta) AgenciaDigito() string {
	d := strings.Split(c.Agencia, "-")
	if len(d) > 1 {
		return d[1]
	}
	return ""
}

func (c Conta) Numero() string {
	n := strings.Split(c.ContaCorrente, "-")
	if len(n) > 0 {
		return n[0]
	}
	return ""
}

func (c Conta) Digito() string {
	d := strings.Split(c.ContaCorrente, "-")
	if len(d) > 1 {
		return d[1]
	}
	return ""
}
