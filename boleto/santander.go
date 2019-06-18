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

func (s Santander) CampoLivre() string {
	convenio := cobranca.Zeros(s.boleto.Conta().Convenio, 7)
	carteira := cobranca.Zeros(s.boleto.Conta().Carteira, 3)
	nossoNumero := cobranca.OnlyNumbers(s.NossoNumero())
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
	n := strconv.FormatInt(s.boleto.Numero(), 10)
	return cobranca.Zeros(n, 12)
}
