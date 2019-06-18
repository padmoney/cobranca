package remessa

import (
	"fmt"
	"strconv"
	"time"

	"github.com/padmoney/cobranca"
)

type SantanderCNAB400 struct {
	conta                cobranca.Conta
	params               Params
	sequencialRegistro   int
	quantidadeDocumentos int
	valorTotalDocumentos float64
}

func NewSantanderCNAB400(c cobranca.Conta, p Params) *SantanderCNAB400 {
	return &SantanderCNAB400{conta: c,
		params: p}
}

func (s SantanderCNAB400) Header() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
		"0",                              // Código do registro = 0
		"1",                              // Código da remessa = 1
		"REMESSA",                        // Literal de transmissão = REMESSA
		"01",                             // Código do serviço = 01
		cobranca.Brancos("COBRANCA", 15), // Literal de serviço = COBRANÇA
		cobranca.Zeros(s.codigoTransmissao(), 20),       // Código de Transmissão (nota 1)
		cobranca.Brancos(s.conta.Beneficiario.Nome, 30), // Nome do Beneficiário
		"033",                     // Código do Banco = 353/033
		"SANTANDER      ",         // Nome do Banco = SANTANDER
		DataFormatada(time.Now()), // Data de Gravação
		cobranca.Zeros("0", 16),   // Zeros
		cobranca.Brancos("", 47),  // Mensagem 1
		cobranca.Brancos("", 47),  // Mensagem 2
		cobranca.Brancos("", 47),  // Mensagem 3
		cobranca.Brancos("", 47),  // Mensagem 4
		cobranca.Brancos("", 47),  // Mensagem 5
		cobranca.Brancos("", 34),  // Brancos
		cobranca.Brancos("", 6),   // Brancos
		cobranca.Zeros("0", 3),    // Número da versão da remessa opcional, se informada, será controlada pelo sistema
		"000001",                  // Número sequencial do registro no arquivo = 000001
	)
}

func (s *SantanderCNAB400) Pagamentos(pagamentos []Pagamento) (lines []string, err error) {
	s.sequencialRegistro = 1
	s.quantidadeDocumentos = 0
	s.valorTotalDocumentos = 0.0
	for _, p := range pagamentos {

		s.quantidadeDocumentos += 1
		s.valorTotalDocumentos += p.Valor()
		s.sequencialRegistro += 1

		indMulta := "0"
		if p.PercentualMulta() > 0.0 {
			indMulta = "4"
		}

		l := fmt.Sprintf("1%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
			s.conta.Beneficiario.TipoInscricao(),
			cobranca.Zeros(s.conta.Beneficiario.GetDocumento(), 14),
			cobranca.Zeros(s.conta.NumeroAgencia(), 4),
			cobranca.Zeros(s.conta.Convenio, 8),
			s.formataNumeroConta(s.conta.Numero(), 8),
			cobranca.Brancos(" ", 25),
			cobranca.Zeros(cobranca.SemMascara(p.NossoNumero()), 8),
			"000000",
			" ",
			indMulta,
			ValorFormatado(p.PercentualMulta(), 2, 2),
			"00",
			cobranca.Zeros("0", 13),
			"    ",
			"000000",                       // Data para cobrança de multa. (Nota 4)
			"5",                            // Código da carteira
			s.comando(p.Comando()),         // Código da ocorrência
			cobranca.Zeros(p.Numero(), 10), // Seu número
			DataFormatada(p.DataVencimento()),
			ValorFormatado(p.Valor(), 11, 2), // 24.7 127 a 139 9(011)v99 Valor do Título
			"033",
			"00000", // Código da agência cobradora do Banco Santander informar somente se carteira for igual a 5, caso contrário, informar zeros.
			cobranca.Zeros(s.conta.GetEspecieTitulo(), 2), // Espécie de documento
			cobranca.Brancos(s.conta.GetAceite(), 1),
			DataFormatada(time.Now()), // Data da emissão do título
			"00",                      // Primeira instrução cobrança
			"00",                      // Segunda instrução cobrança
			ValorFormatado(p.JurosMoraPorDiaAtraso(), 11, 2), // Valor de mora a ser cobrado por dia de atraso
			"000000",                 // Data limite para concessão de desconto
			ValorFormatado(0, 11, 2), // Valor de desconto a ser concedido
			ValorFormatado(0, 8, 5),  // Valor do IOF a ser recolhido pelo Banco para nota de seguro
			ValorFormatado(0, 11, 2), // Valor do abatimento a ser concedido ou valor do segundo desconto. Vide posição 71.
			p.Pagador().TipoInscricao(),
			cobranca.Zeros(p.Pagador().GetDocumento(), 14),
			cobranca.Brancos(p.Pagador().Nome, 40),
			cobranca.Brancos(p.Pagador().Endereco, 40),
			cobranca.Brancos(p.Pagador().Bairro, 12),
			cobranca.Brancos(cobranca.SemMascara(p.Pagador().CEP), 8),
			cobranca.Brancos(p.Pagador().Cidade, 15),
			cobranca.Brancos(p.Pagador().UF, 2),
			cobranca.Brancos(" ", 30), // Nome do sacador ou coobrigado
			" ",
			s.complemento(),
			cobranca.Brancos("", 6),
			"00",
			" ",
			cobranca.Zeros(strconv.Itoa(s.sequencialRegistro), 6), // Sequencial de Registro
		)
		lines = append(lines, l)
	}
	return
}

func (s SantanderCNAB400) Trailler() string {
	q := strconv.Itoa(s.quantidadeDocumentos)
	sr := strconv.Itoa(s.sequencialRegistro + 1)
	return fmt.Sprintf("9%s%s%s%s",
		cobranca.Zeros(q, 6),
		ValorFormatado(s.valorTotalDocumentos, 11, 2),
		cobranca.Zeros("0", 374),
		cobranca.Zeros(sr, 6),
	)
}

func (s SantanderCNAB400) comando(c string) string {
	// Código da ocorrência:
	// 01 = ENTRADA DE TÍTULO
	// 02 = BAIXA DE TÍTULO
	// 04 = CONCESSÃO DE ABATIMENTO
	// 05 = CANCELAMENTO ABATIMENTO
	// 06 = ALTERAÇÃO DE VENCIMENTO
	// 07 = ALT. NÚMERO CONT.BENEFICIÁRIO
	// 08 = ALTERAÇÃO DO SEU NÚMERO
	// 09 = PROTESTAR
	// 18 = SUSTAR PROTESTO (Após início do ciclo de protesto)
	// 98 = NÃO PROTESTAR (Antes do início do ciclo de protesto)
	switch c {
	case ComandoCancelamento:
		return "02"
	default:
		return "01"
	}
}

func (s SantanderCNAB400) contaPadraoNovo() bool {
	return len(s.conta.Numero()) > 8
}

func (s SantanderCNAB400) codigoTransmissao() string {
	return s.params.Get("codigo_transmissao")
}

func (s SantanderCNAB400) complemento() string {
	if s.contaPadraoNovo() {
		n := s.conta.Numero()
		n = n[len(n)-1:]
		return fmt.Sprintf("I%s%s",
			cobranca.Zeros(n, 1),
			cobranca.Zeros(s.conta.Digito(), 1),
		)
	}
	return "   "
}

func (s SantanderCNAB400) formataNumeroConta(c string, q int) string {
	if len(c) > q {
		return c[0:q]
	}
	return cobranca.Zeros(c, q)
}
