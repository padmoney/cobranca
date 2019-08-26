package boleto

import (
	"testing"

	"github.com/padmoney/cobranca"
)

func TestBancoBrasilConvenio4Digitos(t *testing.T) {
	c := contaBancoBrasilFixture("18", "1238")
	b, err := NewBoleto(135.00, Date(2008, 02, 01), 123456, c)

	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	l := "00191238011234564042400061900189137690000013500"
	if b.LinhaDigitavel() != l {
		t.Errorf("Expected '%s' got '%s'", l, b.LinhaDigitavel())
	}

	codBarras := "00191376900000135001238012345640420006190018"
	if b.CodigoBarras() != codBarras {
		t.Errorf("Expected '%s' got '%s'", codBarras, b.CodigoBarras())
	}

	nossoNumero := "12380123456-0"
	if b.NossoNumero() != nossoNumero {
		t.Errorf("Expected '%s' got '%s'", nossoNumero, b.NossoNumero())
	}
}

func TestBancoBrasilConvenio6Digitos(t *testing.T) {
	c := contaBancoBrasilFixture("18", "123879")
	b, err := NewBoleto(135.00, Date(2008, 02, 01), 1234, c)

	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}

	l := "00191238769012344042300061900189237690000013500"
	if b.LinhaDigitavel() != l {
		t.Errorf("Expected '%s' got '%s'", l, b.LinhaDigitavel())
	}

	codBarras := "00192376900000135001238790123440420006190018"
	if b.CodigoBarras() != codBarras {
		t.Errorf("Expected '%s' got '%s'", codBarras, b.CodigoBarras())
	}

	nossoNumero := "12387901234-5"
	if b.NossoNumero() != nossoNumero {
		t.Errorf("Expected '%s' got '%s'", nossoNumero, b.NossoNumero())
	}
}

func TestBancoBrasilConvenio7Digitos(t *testing.T) {
	c := contaBancoBrasilFixture("18", "1238798")
	b, err := NewBoleto(135.00, Date(2008, 02, 03), 7777700168, c)

	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
	codBarras := "00193377100000135000000001238798777770016818"
	if b.CodigoBarras() != codBarras {
		t.Errorf("Expected '%s' got '%s'", codBarras, b.CodigoBarras())
	}

	nossoNumero := "12387987777700168-2"
	if b.NossoNumero() != nossoNumero {
		t.Errorf("Expected '%s' got '%s'", nossoNumero, b.NossoNumero())
	}
}

func TestBancoBrasilConvenio7DigitosCarteira17(t *testing.T) {
	c := contaBancoBrasilFixture("17", "2962087")
	b, err := NewBoleto(5.45, Date(2017, 03, 8), 2, c)

	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
	l := "00190000090296208700900000002170470920000000545"
	if b.LinhaDigitavel() != l {
		t.Errorf("Expected '%s' got '%s'", l, b.LinhaDigitavel())
	}

	codBarras := "00194709200000005450000002962087000000000217"
	if b.CodigoBarras() != codBarras {
		t.Errorf("Expected '%s' got '%s'", codBarras, b.CodigoBarras())
	}

	nossoNumero := "29620870000000002-1"
	if b.NossoNumero() != nossoNumero {
		t.Errorf("Expected '%s' got '%s'", nossoNumero, b.NossoNumero())
	}
}

func TestBancoBrasilConvenio8Digitos(t *testing.T) {
	c := contaBancoBrasilFixture("18", "12387989")
	b, err := NewBoleto(135.00, Date(2008, 02, 02), 7700168, c)

	if err != nil {
		t.Errorf("Expected no error, got '%s'", err.Error())
	}
	l := "00190000090123879890207700168185337700000013500"
	if b.LinhaDigitavel() != l {
		t.Errorf("Expected '%s' got '%s'", l, b.LinhaDigitavel())
	}

	codBarras := "00193377000000135000000001238798900770016818"
	if b.CodigoBarras() != codBarras {
		t.Errorf("Expected '%s' got '%s'", codBarras, b.CodigoBarras())
	}

	nossoNumero := "12387989007700168-7"
	if b.NossoNumero() != nossoNumero {
		t.Errorf("Expected '%s' got '%s'", nossoNumero, b.NossoNumero())
	}
}

func TestBancoBrasilConvenioInvalido(t *testing.T) {
	c := contaBancoBrasilFixture("18", "123879890")
	_, err := NewBoleto(135.00, Date(2008, 02, 02), 7700168, c)
	if err == nil {
		t.Errorf("Expected error, got no error")
	}
	if err.Error() != errTipoConvenioNaoImplementado {
		t.Errorf("Expected '%s' got '%s'", errTipoConvenioNaoImplementado, err.Error())
	}
}

func contaBancoBrasilFixture(carteira, convenio string) cobranca.Conta {
	return NewConta(cobranca.CodigoBancoBrasil, "4042", "61900", carteira, convenio)
}
