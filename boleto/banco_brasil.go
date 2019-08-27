package boleto

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/padmoney/cobranca"
)

type BancoBrasil struct {
	boleto Boleto
}

func NewBancoBrasil(boleto Boleto) BancoBrasil {
	return BancoBrasil{boleto: boleto}
}

func (bb BancoBrasil) CampoLivre() (string, error) {
	convenio := bb.boleto.Conta().Convenio
	carteira := cobranca.Zeros(bb.boleto.Conta().Carteira, 2)
	nossoNumero, err := bb.nossoNumeroSemDV()
	if err != nil {
		return "", err
	}
	var campoLivre string
	switch len(convenio) {
	case 4, 6:
		campoLivre = fmt.Sprintf("%s%s%s%s",
			cobranca.Zeros(nossoNumero, 11),
			cobranca.Zeros(bb.boleto.Conta().Agencia, 4),
			cobranca.Zeros(bb.boleto.Conta().Numero(), 8),
			carteira)
	case 7, 8:
		campoLivre = fmt.Sprintf("000000%s%s", cobranca.Zeros(nossoNumero, 17), carteira)
	default:
		return "", errors.New(errTipoConvenioNaoImplementado)
	}
	return campoLivre, nil
}

func (bb BancoBrasil) NossoNumero() (string, error) {
	cb, err := bb.nossoNumeroSemDV()
	if err != nil {
		return "", err
	}
	dv, err := Modulo11{}.Calculate(cb)
	if err != nil {
		return "", err
	}
	return cb + "-" + dv, nil
}

func (bb BancoBrasil) nossoNumeroSemDV() (string, error) {
	var quantidade int
	convenio := bb.boleto.Conta().Convenio
	switch len(convenio) {
	case 4:
		quantidade = 7
	case 7:
		quantidade = 10
	case 6:
		quantidade = 5
	case 8:
		quantidade = 9
	default:
		return "", errors.New(errTipoConvenioNaoImplementado)
	}
	n := strconv.FormatInt(bb.boleto.Numero(), 10)
	var ultimosDigitos string
	if len(n)-quantidade > 0 {
		ultimosDigitos = n[len(n)-quantidade:]
	} else {
		ultimosDigitos = n
	}
	numero := cobranca.Zeros(ultimosDigitos, quantidade)
	return fmt.Sprintf("%s%s", convenio, numero), nil
}
