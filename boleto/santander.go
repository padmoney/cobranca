package boleto

import (
	"strconv"
)

type Santander struct {
	boleto Boleto
}

func NewSantander(boleto Boleto) Santander {
	return Santander{boleto: boleto}
}

func (s Santander) CampoLivre() string {
	convenio := Zeros(s.boleto.Conta().Convenio, 7)
	carteira := Zeros(s.boleto.Conta().Carteira, 3)
	nossoNumero := OnlyNumbers(s.NossoNumero())
	return "9" + convenio + nossoNumero + "0" + carteira
}

func (s Santander) NossoNumero() string {
	cb := s.codigoBarrasSemDV()
	dv, err := Modulo11{}.Calculate(cb)
	if err != nil {
		return ""
	}
	return cb + "-" + dv
}

func (s Santander) codigoBarrasSemDV() string {
	n := strconv.Itoa(s.boleto.Numero())
	return Zeros(n, 12)
}
