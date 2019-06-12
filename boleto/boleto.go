package boleto

import (
	"errors"
	"time"

	"github.com/padmoney/cobranca"
)

type BoletoGenerator interface {
	NossoNumero() string
	CampoLivre() string
}

type Boleto struct {
	valor             float64
	dataDocumento     time.Time
	dataProcessamento time.Time
	dataVencimento    time.Time
	numero            int64
	nossoNumero       string
	campoLivre        string
	conta             cobranca.Conta
	pagador           Pagador
	avalista          Avalista
	codigoBarras      string
	linhaDigitavel    string
}

func NewBoleto(valor float64, dataVencimento time.Time, numero int64, conta cobranca.Conta) (Boleto, error) {
	boleto := Boleto{
		valor:          valor,
		dataVencimento: dataVencimento,
		numero:         numero,
		conta:          conta,
	}

	var bg BoletoGenerator
	switch conta.Banco {
	case cobranca.CodigoSantander:
		bg = NewSantander(boleto)
	default:
		return boleto, errors.New(cobranca.BancoNaoSuportado)
	}
	boleto.nossoNumero = bg.NossoNumero()
	boleto.campoLivre = bg.CampoLivre()
	return boleto, nil
}

func (b Boleto) CodigoBarras() string {
	if b.codigoBarras == "" {
		b.codigoBarras = CodigoBarras{Banco: b.conta.Banco,
			Valor:          b.valor,
			DataVencimento: b.dataVencimento,
			CampoLivre:     b.campoLivre}.String()
	}
	return b.codigoBarras
}

func (b Boleto) Conta() cobranca.Conta {
	return b.conta
}

func (b Boleto) DataDocumento() time.Time {
	return time.Now().Local()
}

func (b Boleto) DataVencimento() time.Time {
	return b.dataVencimento
}

func (b Boleto) LinhaDigitavel() string {
	if b.linhaDigitavel == "" {
		cb := b.CodigoBarras()
		b.linhaDigitavel, _ = LinhaDigitavel{}.Calculate(cb)
	}
	return b.linhaDigitavel
}

func (b Boleto) LocalPagamento() string {
	return "Pagável em qualquer banco até o vencimento."
}

func (b Boleto) NossoNumero() string {
	return b.nossoNumero
}

func (b Boleto) Numero() int64 {
	return b.numero
}

func (b Boleto) Valor() float64 {
	return b.valor
}
