package boleto

import (
	"errors"

	"github.com/padmoney/cobranca"
)

type nossoNumero struct {
	banco  string
	numero string
}

func NossoNumero(numero string) *nossoNumero {
	return &nossoNumero{numero: numero}
}

func (nn *nossoNumero) BancoBrasil() *nossoNumero {
	nn.banco = cobranca.CodigoBancoBrasil
	return nn
}

func (nn nossoNumero) DV() (string, error) {
	switch nn.banco {
	case cobranca.CodigoBancoBrasil:
		m := map[int]string{10: "X"}
		return Modulo11{Mapping: m}.Calculate(nn.numero)
	}
	return "", errors.New(cobranca.BancoNaoSuportado)
}
