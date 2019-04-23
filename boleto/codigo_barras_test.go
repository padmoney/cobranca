package boleto

import (
	"testing"
	"time"

	"github.com/padmoney/cobranca"
)

type CodigoBarrasFixture struct {
	banco          string
	valor          float64
	dataVencimento time.Time
	campoLivre     string
}

func TestCodigoBarras(t *testing.T) {
	m := map[string]CodigoBarrasFixture{
		"00197000000000005450000002962087000000000117": CodigoBarrasFixture{banco: cobranca.CodigoBancoBrasil,
			valor:          5.45,
			dataVencimento: Date(1997, 10, 07),
			campoLivre:     "0000002962087000000000117"},
		"00199000000000042090000002962087000000000217": CodigoBarrasFixture{banco: cobranca.CodigoBancoBrasil,
			valor:          42.09,
			dataVencimento: Date(1997, 10, 07),
			campoLivre:     "0000002962087000000000217"},
		"00198000000000176720000002962087000000000317": CodigoBarrasFixture{banco: cobranca.CodigoBancoBrasil,
			valor:          176.72,
			dataVencimento: Date(1997, 10, 07),
			campoLivre:     "0000002962087000000000317"},
		"00199000000001540360000002962087000000000417": CodigoBarrasFixture{banco: cobranca.CodigoBancoBrasil,
			valor:          1540.36,
			dataVencimento: Date(1997, 10, 07),
			campoLivre:     "0000002962087000000000417"},
		"03394740000012345679897033512345678901230101": CodigoBarrasFixture{banco: cobranca.CodigoSantander,
			valor:          12345.67,
			dataVencimento: Date(2018, 01, 10),
			campoLivre:     "9897033512345678901230101"}}

	for expected, v := range m {
		got := CodigoBarras{Banco: string(v.banco),
			Valor:          v.valor,
			DataVencimento: v.dataVencimento,
			CampoLivre:     v.campoLivre}.String()
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestCodigoBarrasMoeda(t *testing.T) {
	expected := "X"
	cb := CodigoBarras{}
	cb.SetMoeda(expected)
	got := cb.Moeda()
	if got != expected {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}

func TestFatorVencimento(t *testing.T) {
	m := map[string]time.Time{
		"0020": Date(1997, 10, 27),
		"1000": Date(2000, 07, 03),
		"1001": Date(2000, 07, 04),
		"1002": Date(2000, 07, 05),
		"1667": Date(2002, 05, 01),
		"2046": Date(2003, 05, 15),
		"3737": Date(2007, 12, 31),
		"4789": Date(2010, 11, 17),
		"7400": Date(2018, 01, 10),
		"9999": Date(2025, 02, 21),
		"9474": Date(2023, 9, 15),
		"7561": Date(2018, 6, 20)}

	for expected, v := range m {
		got := FatorVencimento(v)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}
