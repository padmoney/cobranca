package boleto

import (
	"testing"
	"time"

	"github.com/padmoney/cobranca"
)

func TestNewBoleto(t *testing.T) {
	dateFormat := "2006-01-02"
	localPagto := "Pagável em qualquer banco até o vencimento."
	venc := Date(2003, 05, 15)
	today := time.Now().Local().Format(dateFormat)
	valor := 273.71

	c := NewConta(cobranca.CodigoSantander, "4042", "61900", "101", "0282033")

	b, _ := NewBoleto(valor, venc, 1984, c)

	if b.LocalPagamento() != localPagto {
		t.Errorf("Expected '%s' got '%s'", localPagto, b.LocalPagamento())
	}

	dataDoc := b.DataDocumento().Format(dateFormat)
	if dataDoc != today {
		t.Errorf("Expected '%s' got '%s'", today, dataDoc)
	}

	if b.DataVencimento() != venc {
		t.Errorf("Expected '%s' got '%s'", venc, b.DataVencimento())
	}

	if b.Valor() != valor {
		t.Errorf("Expected '%f' got '%f'", valor, b.Valor())
	}

	nossoNumero := "000000001984-4"
	if b.NossoNumero() != nossoNumero {
		t.Errorf("Expected '%s' got '%s'", nossoNumero, b.NossoNumero())
	}

	linhaDigitavel := "03399028270330000000101984401016920460000027371"
	if b.LinhaDigitavel() != linhaDigitavel {
		t.Errorf("Expected '%s' got '%s'", linhaDigitavel, b.LinhaDigitavel())
	}

	codBarras := "03399204600000273719028203300000000198440101"
	if b.CodigoBarras() != codBarras {
		t.Errorf("Expected '%s' got '%s'", codBarras, b.CodigoBarras())
	}
}

func TestBoletoBancoNaoSuportado(t *testing.T) {
	c := NewConta("999", "4042", "61900", "101", "0282033")
	_, err := NewBoleto(1, time.Now(), 1984, c)

	if err == nil {
		t.Errorf("Should'n be nil")
	}

	if err.Error() != cobranca.BancoNaoSuportado {
		t.Errorf("Expected '%s' got '%s'", cobranca.BancoNaoSuportado, err.Error())
	}
}
