package retorno

import (
	"testing"

	"github.com/padmoney/cobranca"
)

var fileBB = `02RETORNO01COBRANCA       12343000153990000000XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX001BANCO DO BRASIL0309190000578                      000000210987654321  1234567                                                                                                                                                                                                                                              000001
70000000000000000123430001539901234567                         2994806000006821770000301AI 02700000000000 17060309190000068217                    060919000000000399023714430010409190000128000000000000000000000000000000000000000000000000000000000000000000000000003990000000000000000000000000000000000000000000000000386220000000000000          0000000000000000000000000000000000000000000000001010000002
9201001          000997910000039816309000000249          000000000000000000000000000000          000000000000000000000000000000          000000000000000000000000000000                                                  000000000000000000000000000000                                                                                                                                                   000003`

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

func TestBancoBrasilCNAB400Conteudo(t *testing.T) {
	r, err := New(cobranca.CodigoBancoBrasil, CNAB400)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	registros, err := r.ReadFile([]byte(fileBB))
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	if len(registros) != 1 {
		t.Errorf("Expected '%d', got '%d'", 3, len(registros))
	}

	r1 := registros[0]
	nn1 := "29948060000068217-5"
	if r1.NossoNumero != nn1 {
		t.Errorf("Expected '%s', got '%s'", nn1, r1.NossoNumero)
	}
	v1 := 39.9
	if r1.Valor != v1 {
		t.Errorf("Expected '%f', got '%f'", v1, r1.Valor)
	}
}
