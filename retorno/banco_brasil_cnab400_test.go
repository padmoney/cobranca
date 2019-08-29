package retorno

import (
	"testing"

	"github.com/padmoney/cobranca"
)

func TestBancoBrasilCNAB400(t *testing.T) {
	r, err := New(cobranca.CodigoBancoBrasil, CNAB400)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	registros, err := r.Read("arquivos/CBR6437172808201910804.ret")
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	if len(registros) != 3 {
		t.Errorf("Expected '%d', got '%d'", 3, len(registros))
	}

	r1 := registros[0]
	nn1 := "12345670000004600-5"
	if r1.NossoNumero != nn1 {
		t.Errorf("Expected '%s', got '%s'", nn1, r1.NossoNumero)
	}
	v1 := 39.9
	if r1.Valor != v1 {
		t.Errorf("Expected '%f', got '%f'", v1, r1.Valor)
	}
}
