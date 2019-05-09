package boleto

import (
	"errors"
	"time"

	"github.com/padmoney/cobranca"
)

const (
	bancoNaoSuportado = "Banco não suportado."
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
	conta             Conta
	pagador           Pagador
	avalista          Avalista
}

func NewBoleto(valor float64, dataVencimento time.Time, numero int64, conta Conta) (Boleto, error) {
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
		return boleto, errors.New(bancoNaoSuportado)
	}
	boleto.nossoNumero = bg.NossoNumero()
	boleto.campoLivre = bg.CampoLivre()
	return boleto, nil
}

func (b Boleto) CodigoBarras() string {
	return CodigoBarras{Banco: b.conta.Banco,
		Valor:          b.valor,
		DataVencimento: b.dataVencimento,
		CampoLivre:     b.campoLivre}.String()
}

func (b Boleto) Conta() Conta {
	return b.conta
}

func (b Boleto) DataDocumento() time.Time {
	return time.Now().Local()
}

func (b Boleto) DataVencimento() time.Time {
	return b.dataVencimento
}

func (b Boleto) LinhaDigitavel() string {
	cb := b.CodigoBarras()
	l, _ := LinhaDigitavel{}.Calculate(cb)
	return l
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
