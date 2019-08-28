package retorno

import (
	"testing"

	"github.com/padmoney/cobranca"
)

func TestNewRetorno(t *testing.T) {
	_, err := New(cobranca.CodigoBancoBrasil, CNAB400)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}

	_, err = New(cobranca.CodigoSantander, CNAB400)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
}

func TestNewRetornoBancoNaoSuportado(t *testing.T) {
	_, err := New("", CNAB400)
	if err.Error() != cobranca.BancoNaoSuportado {
		t.Errorf("Expected %s, got %s", cobranca.BancoNaoSuportado, err.Error())
	}

	_, err = New("999", CNAB400)
	if err.Error() != cobranca.BancoNaoSuportado {
		t.Errorf("Expected %s, got %s", cobranca.BancoNaoSuportado, err.Error())
	}
}

func TestNewRetornoLayoutNaoSuportado(t *testing.T) {
	_, err := New(cobranca.CodigoBancoBrasil, CNAB240)
	if err.Error() != layoutRetornoNaoSuportado {
		t.Errorf("Expected %s, got %s", layoutRetornoNaoSuportado, err.Error())
	}

	_, err = New(cobranca.CodigoSantander, CNAB240)
	if err.Error() != layoutRetornoNaoSuportado {
		t.Errorf("Expected %s, got %s", layoutRetornoNaoSuportado, err.Error())
	}
}

func TestNewRetornoArquivoRetornoNaoEncontrado(t *testing.T) {
	r, err := New(cobranca.CodigoBancoBrasil, CNAB400)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	_, err = r.Read("")
	if err.Error() != arquivoRetornoNaoEncontrado {
		t.Errorf("Expected %s, got %s", arquivoRetornoNaoEncontrado, err.Error())
	}

	_, err = New(cobranca.CodigoSantander, CNAB400)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	_, err = r.Read("")
	if err.Error() != arquivoRetornoNaoEncontrado {
		t.Errorf("Expected %s, got %s", arquivoRetornoNaoEncontrado, err.Error())
	}
}
