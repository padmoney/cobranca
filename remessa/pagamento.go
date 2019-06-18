package remessa

import (
	"fmt"
	"time"

	"github.com/padmoney/cobranca"
)

type Pagamento struct {
	comando             string
	valor               float64
	dataVencimento      time.Time
	numero              string
	nossoNumero         string
	diasProtesto        int
	percentualMoraAoMes float64
	percentualMulta     float64
	pagador             cobranca.Pagador
	avalista            cobranca.Avalista
}

// NewPagamento retorna um pagamento
// comando: baixa, registro
func NewPagamento(
	comando string,
	valor float64,
	dataVencimento time.Time,
	numero string,
	nossoNumero string,
	diasProtesto int,
	percentualMoraAoMes float64,
	percentualMulta float64,
	pagador cobranca.Pagador,
	avalista cobranca.Avalista) Pagamento {
	return Pagamento{
		comando:             comando,
		valor:               valor,
		dataVencimento:      dataVencimento,
		numero:              numero,
		nossoNumero:         nossoNumero,
		diasProtesto:        diasProtesto,
		percentualMoraAoMes: percentualMoraAoMes,
		percentualMulta:     percentualMulta,
		pagador:             pagador,
		avalista:            avalista,
	}
}

/**
 * Mensagem Avalista
 * Para CNPJ
 *  Posição 352 à 372 - Preencher com o nome do Sacador/Avalista.
 *  Posição 373 - Preencher com "espaço"
 *  Posição 374 à 377 - Preencher com o literal "CNPJ"
 *  Posição 378 à 391 - Preencher com o número do CNPJ do Sacador/Avalista
 * Para CPF
 *  Posição 352 à 376 - Preencher com o nome do Sacador/Avalista
 *  Posição 377 - Preencher com "espaço"
 *  Posição 378 à 380 - Preencher com o literal "CPF"
 *  Posição 381 à 391 - Preencher com o número do CPF do Sacador/Avalista
 *
 * @return Pessoa Retorna o avalista
 */
func (p Pagamento) MensagemAvalista() string {
	doc := cobranca.SemMascara(p.avalista.Documento)
	var brancos int
	var tipoDoc string
	switch len(doc) {
	case 11:
		tipoDoc = "CPF"
		brancos = 26
	case 14:
		tipoDoc = "CNPJ"
		brancos = 22
	default:
		return p.avalista.Nome
	}
	return fmt.Sprintf("%s%s%s",
		cobranca.Brancos(p.avalista.Nome, brancos),
		tipoDoc,
		doc)
}

func (p Pagamento) JurosMoraPorDiaAtraso() float64 {
	percJurosMoraDia := (p.percentualMoraAoMes / 100.0) / 30.0
	return percJurosMoraDia * p.valor
}

func (p Pagamento) Comando() string {
	return p.comando
}

func (p Pagamento) Valor() float64 {
	return p.valor
}

func (p Pagamento) DataVencimento() time.Time {
	return p.dataVencimento
}

func (p Pagamento) Numero() string {
	return p.numero
}

func (p Pagamento) NossoNumero() string {
	return p.nossoNumero
}

func (p Pagamento) DiasProtesto() int {
	return p.diasProtesto
}

func (p Pagamento) PercentualMoraAoMes() float64 {
	return p.percentualMoraAoMes
}

func (p Pagamento) PercentualMulta() float64 {
	return p.percentualMulta
}

func (p Pagamento) Pagador() cobranca.Pagador {
	return p.pagador
}

func (p Pagamento) Avalista() cobranca.Avalista {
	return p.avalista
}
