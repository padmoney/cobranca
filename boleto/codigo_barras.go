package boleto

import (
	"math"
	"strconv"
	"time"

	"github.com/padmoney/cobranca"
)

type CodigoBarras struct {
	Banco          string
	Valor          float64
	CodigoMoeda    string
	DataVencimento time.Time
	CampoLivre     string
}

func (cb CodigoBarras) String() string {
	banco := cobranca.StrPad(cb.Banco, 3, "0", cobranca.StrPadLeft)
	val := math.Round(cb.Valor * 100)
	valStr := strconv.Itoa(int(val))
	valStr = cobranca.Zeros(valStr, 10)
	fatorVencimento := FatorVencimento(cb.DataVencimento)

	s := banco + cb.Moeda() + fatorVencimento + valStr + cb.CampoLivre
	dv, _ := Modulo11{Mapping: map[int]string{0: "1"}}.Calculate(s)
	return banco + cb.Moeda() + dv + fatorVencimento + valStr + cb.CampoLivre
}

func (cb *CodigoBarras) SetMoeda(m string) {
	cb.CodigoMoeda = m
}

func (cb CodigoBarras) Moeda() string {
	if cb.CodigoMoeda == "" {
		return "9"
	}
	return cb.CodigoMoeda
}

func FatorVencimento(d time.Time) string {
	base := Date(1997, 10, 07)
	delta := d.Sub(base)
	days := delta.Hours() / 24
	s := strconv.Itoa(int(days))
	return cobranca.Zeros(s, 4)
}
