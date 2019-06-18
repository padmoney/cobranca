package remessa

import (
	"testing"
	"time"

	"github.com/padmoney/cobranca"
)

func TestCNAB400(t *testing.T) {

	expectedHeader := "01REMESSA01COBRANCA       17777751042700080112ACME CORPORATION              033SANTANDER      " +
		DataFormatada(time.Now()) +
		"0000000000000000                    " +
		cobranca.Brancos("", 255) +
		"000000001"
	expectedDetalheAntesData := "1029999999900019140420123456713003758                         00000123000000 40100000000000000000    000000501000000012317071900000000001990330000001A"
	expectedDetalheDepoisData := "000000000000000000000000000000000000000000000000000000000000000200000000000191NOME DO SACADO                          ENDERECO DO SACADO                      BAIRRO      29315732CACHOEIRO DE ITES                               I90      00 000002"

	expectedDetalhe := expectedDetalheAntesData + DataFormatada(time.Now()) + expectedDetalheDepoisData

	expectedTrailer := "9000001000000000019900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003"

	conta := fixtureConta("18", "1234567")
	params := fixtureParams()
	remessa, err := New("CNAB400", conta, params, 1)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}
	p := fixturePagamento(1.99, time.Now().AddDate(0, 1, 0), "123", "123")

	remessa.AddPagamento(p)

	linhas, err := remessa.Strings()
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}

	if got := len(linhas); got != 3 {
		t.Errorf("Expected '%d', got '%d'", 3, got)
	}

	// header
	if got := linhas[0]; got != expectedHeader {
		t.Errorf("header: Expected '%s', got '%s'", expectedHeader, got)
	}
	if got := len(linhas[0]); got != 400 {
		t.Errorf("header: Expected '%d', got '%d'", 400, got)
	}

	// detalhe
	if got := linhas[1]; got != expectedDetalhe {
		t.Errorf("detail: Expected '%s', got '%s'", expectedDetalhe, got)
	}
	if got := len(linhas[1]); got != 400 {
		t.Errorf("detail: Expected '%d', got '%d'", 400, got)
	}

	// trailer
	if got := linhas[2]; got != expectedTrailer {
		t.Errorf("trailer: Expected '%s', got '%s'", expectedTrailer, got)
	}
	if got := len(linhas[2]); got != 400 {
		t.Errorf("trailer: Expected '%d', got '%d'", 400, got)
	}
	//        $this->assertEquals('REMESSA_1.rem', $remessa->nomeArquivo(), 'Nome do arquivo');

}

func fixtureConta(carteira, convenio string) cobranca.Conta {
	beneficiario := cobranca.Beneficiario{
		Nome:      "ACME Corporation",
		Documento: "99999999000191",
	}
	return cobranca.Conta{
		Banco:         "033",
		Agencia:       "4042",
		ContaCorrente: "130037589",
		Carteira:      carteira,
		Convenio:      convenio,
		Beneficiario:  beneficiario,
		Aceite:        true,
	}
}

func fixturePagamento(valor float64, dataVencimento time.Time, numero, nossoNumero string) Pagamento {
	avalista := cobranca.Avalista{Nome: "Nome do Avalista", Documento: "335.176.000-08"}
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
	/*

	   private function novoPagamento($valor, $data_vencimento, $numero, $nosso_numero = null,
	       $dias_protesto = 30,
	       $percentual_mora_ao_mes = 1,
	       $percentual_multa = 1,
	       $avalista = []
	   ) {
	       empty($nosso_numero) && $nosso_numero = $numero;
	       $pagador = [
	       ];
	       return new Pagamento('registro', $valor, $data_vencimento, $numero, $nosso_numero,
	           $dias_protesto,
	           $percentual_mora_ao_mes,
	           $percentual_multa,
	           $pagador);
	   }
	*/
}

func fixtureParams() Params {
	p := NewParams()
	p.Add("codigo_transmissao", "17777751042700080112")
	return p
}
