package boleto

import (
	"testing"

	"github.com/padmoney/cobranca"
)

func TestBoletoSantander(t *testing.T) {
	c := contaSantanderFixture("101", "0282033")
	b, _ := NewBoleto(34.80, Date(2018, 06, 22), 1984, c)

	l := "03399028270330000000101984401016275630000003480"
	if b.LinhaDigitavel() != l {
		t.Errorf("Expected '%s' got '%s'", l, b.LinhaDigitavel())
	}

	codBarras := "03392756300000034809028203300000000198440101"
	if b.CodigoBarras() != codBarras {
		t.Errorf("Expected '%s' got '%s'", codBarras, b.CodigoBarras())
	}
}

func TestBoletoSantanderSemRegistro(t *testing.T) {
	c := contaSantanderFixture("102", "7041160")
	b, _ := NewBoleto(11.5, Date(2017, 8, 10), 701650257, c)

	l := "03399704101600007016550257101027272470000001150"
	if b.LinhaDigitavel() != l {
		t.Errorf("Expected '%s' got '%s'", l, b.LinhaDigitavel())
	}
}

func contaSantanderFixture(carteira, convenio string) Conta {
	return NewConta(cobranca.CodigoSantander, "4042", "61900", carteira, convenio)
}
