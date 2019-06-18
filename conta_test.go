package cobranca

import (
	"testing"
)

func TestNewConta(t *testing.T) {
	agencia := "1234"
	contaCorrente := "567890"
	carteira := "01"
	convenio := "123456789"
	beneficiario := Beneficiario{}

	c := NewConta(CodigoBancoBrasil, agencia, contaCorrente, carteira, convenio, beneficiario)

	if c.Banco != CodigoBancoBrasil {
		t.Errorf("Expected '%s' got '%s'", CodigoBancoBrasil, c.Banco)
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
