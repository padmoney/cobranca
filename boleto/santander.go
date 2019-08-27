package boleto

import (
	"strconv"

	"github.com/padmoney/cobranca"
)

type Santander struct {
	boleto Boleto
}

func NewSantander(boleto Boleto) Santander {
	return Santander{boleto: boleto}
}

func (s Santander) CampoLivre() (string, error) {
	convenio := cobranca.Zeros(s.boleto.Conta().Convenio, 7)
	carteira := cobranca.Zeros(s.boleto.Conta().Carteira, 3)
	nossoNumero, err := s.NossoNumero()
	if err != nil {
		return "", nil
	}
	nossoNumero = cobranca.OnlyNumbers(nossoNumero)
	return "9" + convenio + nossoNumero + "0" + carteira, nil
}

func (s Santander) NossoNumero() (string, error) {
	cb := s.nossoNumeroSemDV()
	dv, err := Modulo11{}.Calculate(cb)
	if err != nil {
		return "", nil
	}
	return cb + "-" + dv, nil
}

func (s Santander) nossoNumeroSemDV() string {
	n := strconv.FormatInt(s.boleto.Numero(), 10)
	return cobranca.Zeros(n, 12)
}
