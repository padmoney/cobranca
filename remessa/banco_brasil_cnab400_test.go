package remessa

import (
	"fmt"
	"testing"
	"time"

	"github.com/padmoney/cobranca"
)

func TestBancoBrasilCNAB400(t *testing.T) {
	expectedHeader := fmt.Sprintf("%s%s%s%s%s",
		"01REMESSA01COBRANCA       40428000619000000000ACME CORPORATION              001BANCODOBRASIL  ",
		DataFormatada(time.Now()),
		"0000001                      1234567",
		cobranca.Brancos("", 258),
		"000001",
	)

	expectedDetalhe := fmt.Sprintf("%s%s%s%s%s%s",
		"70299999999000191404280006190001234567                         123456700000001230000       0190000000     ",
		"1801000000012316021700000000001990010000 01A",
		DataFormatada(time.Now()),
		"000000000000000000000000000000000000000000000000000000000000000200000000000191",
		"NOME DO SACADO                          ENDERECO DO SACADO                      BAIRRO      29315732CACHOEIRO DE ITES",
		"                                        30 000002",
	)

	expectedDetalheMulta := "59921602170000000000100                                                                                                                                                                                                                                                                                                                                                                                    000003"

	expectedDetalheComAvalista := fmt.Sprintf("%s%s%s%s%s%s",
		"70299999999000191404280006190001234567                         123456700000001240000   A   0190000000     ",
		"1801000000012416031700000000421990010000 01A",
		DataFormatada(time.Now()),
		"000000000000000140000000000000000000000000000000000000000000000200000000000191",
		"NOME DO SACADO                          ENDERECO DO SACADO                      BAIRRO      29315732CACHOEIRO DE ITES",
		"NOME DO AVALISTA          CPF3351760000830 000004",
	)

	expectedTrailer := fmt.Sprintf("9%s000006", cobranca.Brancos("", 393))

	conta := fixtureContaBancoBrasil("18", "1234567")
	params := fixtureParamsBB("019")
	remessa, err := New("CNAB400", conta, params, 1)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	p1 := fixturePagamentoSemAvalista(1.99, cobranca.Date(2017, 2, 16), "123", "")
	remessa.AddPagamento(p1)
	p2 := fixturePagamento(421.99, cobranca.Date(2017, 3, 16), "124", "")
	remessa.AddPagamento(p2)
	linhas, err := remessa.Linhas()
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}

	if got := len(linhas); got != 6 {
		t.Errorf("Expected '%d', got '%d'", 6, got)
	}

	header := linhas[0]
	if header != expectedHeader {
		t.Errorf("header: Expected '%s', got '%s'", expectedHeader, header)
	}
	if got := len(header); got != 400 {
		t.Errorf("header: Expected '%d', got '%d'", 400, got)
	}

	detalhe := linhas[1]
	if detalhe != expectedDetalhe {
		t.Errorf("detalhe: Expected '%s', got '%s'", expectedDetalhe, detalhe)
	}
	if got := len(header); got != 400 {
		t.Errorf("detalhe: Expected '%d', got '%d'", 400, got)
	}

	detalheMulta := linhas[2]
	if detalheMulta != expectedDetalheMulta {
		t.Errorf("detalheMulta: Expected '%s', got '%s'", expectedDetalheMulta, detalheMulta)
	}
	if got := len(header); got != 400 {
		t.Errorf("detalheMulta: Expected '%d', got '%d'", 400, got)
	}

	detalheComAvalista := linhas[3]
	if detalheComAvalista != expectedDetalheComAvalista {
		t.Errorf("detalhe com avalista: Expected '%s', got '%s'", expectedDetalheComAvalista, detalheComAvalista)
	}
	if got := len(header); got != 400 {
		t.Errorf("detalhe com avalista: Expected '%d', got '%d'", 400, got)
	}

	trailer := linhas[5]
	if trailer != expectedTrailer {
		t.Errorf("trailer: Expected '%s', got '%s'", expectedTrailer, trailer)
	}
	if got := len(header); got != 400 {
		t.Errorf("trailer: Expected '%d', got '%d'", 400, got)
	}
}

func fixtureContaBancoBrasil(carteira, convenio string) cobranca.Conta {
	beneficiario := cobranca.Beneficiario{
		Nome:      "ACME Corporation",
		Documento: "99.999.999/0001-91",
	}
	return cobranca.Conta{
		Banco:         "001",
		Agencia:       "4042-8",
		ContaCorrente: "61900-0",
		Carteira:      carteira,
		Convenio:      convenio,
		Beneficiario:  beneficiario,
		Aceite:        true,
	}
}

func fixturePagamentoSemAvalista(valor float64, dataVencimento time.Time, numero, nossoNumero string) Pagamento {
	avalista := cobranca.Avalista{}
	pagador := cobranca.Pagador{
		Documento: "00000000000191",
		Nome:      "Nome do sacado",
		Endereco:  "Endereço do sacado",
		Bairro:    "Bairro",
		CEP:       "29315-732",
		Cidade:    "Cachoeiro de Itapemirim",
		UF:        "ES",
	}
	return NewPagamento("registro", valor, dataVencimento, numero, nossoNumero, 30, 1.0, 1.0, pagador, avalista)
}

func fixtureParamsBB(variacao string) Params {
	p := NewParams()
	p.Add("variacao", variacao)
	return p
}
