package remessa

import (
	"testing"

	"github.com/padmoney/cobranca"
)

func TestMensagemAvalista(t *testing.T) {
	a := cobranca.Avalista{Nome: "Nome do Avalista",
		Documento: "280.834.810-09"}
	p := Pagamento{avalista: a}
	expected := "NOME DO AVALISTA          CPF28083481009"
	if got := p.MensagemAvalista(); got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
