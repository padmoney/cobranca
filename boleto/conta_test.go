package boleto

import (
	"testing"

	"github.com/padmoney/cobranca"
)

func TestNewConta(t *testing.T) {
	agencia := "1234"
	contaCorrente := "567890"
	carteira := "01"
	convenio := "123456789"

	c := NewConta(cobranca.CodigoBancoBrasil, agencia, contaCorrente, carteira, convenio)

	if c.Banco != cobranca.CodigoBancoBrasil {
		t.Errorf("Expected '%s' got '%s'", cobranca.CodigoBancoBrasil, c.Banco)
	}

	if c.Agencia != agencia {
		t.Errorf("Expected '%s' got '%s'", agencia, c.Agencia)
	}

	if c.ContaCorrente != contaCorrente {
		t.Errorf("Expected '%s' got '%s'", contaCorrente, c.ContaCorrente)
	}

	if c.Carteira != carteira {
		t.Errorf("Expected '%s' got '%s'", carteira, c.Carteira)
	}

	if c.Convenio != convenio {
		t.Errorf("Expected '%s' got '%s'", convenio, c.Convenio)
	}
}