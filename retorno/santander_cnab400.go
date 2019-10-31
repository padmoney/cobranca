package retorno

import (
	"errors"
	"strconv"
	"strings"

	"github.com/padmoney/cobranca"
)

type SantanderCNAB400 struct {
}

func NewSantanderCNAB400() SantanderCNAB400 {
	return SantanderCNAB400{}
}

func (b SantanderCNAB400) LerLinha(linha string) (r Registro, err error) {
	if linha == "" {
		return
	}
	r.ID = linha[0:1]
	switch r.ID {
	case "0":
		if linha[0:9] != "02RETORNO" {
			err = errors.New("Tipo de Operação inválido.")
			return
		}
		if linha[9:19] != "01COBRANCA" {
			err = errors.New("Tipo de Serviço inválido.")
			return
		}
		r.ID = ""
	case "1":
		r, err = b.lerDetalhe(r, linha)
	case "9":
		r.ID = ""
	}
	return
}

func (b SantanderCNAB400) lerDetalhe(reg Registro, s string) (r Registro, err error) {
	r = reg

	r.TipoCobranca = s[107:108]
	comando := s[108:110]
	r.TipoOcorrencia, r.Ocorrencia = b.ocorrencia(comando)
	if r.TipoOcorrencia == liquidado {
		r.DataLiquidacao = parseDate(s[110:116])
	}
	r.Numero = strings.Replace(s[116:126], " ", "", -1)
	r.NossoNumero = b.nossoNumero(s[126:134])
	r.DataVencimento = parseDate(s[146:152])
	r.Valor = parseFloat(s[152:165])
	r.EspecieTitulo = s[173:175]
	r.ValorTarifa = parseFloat(s[175:188])
	r.OutrasDespesas = parseFloat(s[188:201])
	r.JurosAtraso = parseFloat(s[201:214])
	r.TaxaIOF = parseFloat(s[214:217])
	r.ValorAbatimento = parseFloat(s[227:240])
	r.DescontoConcedido = parseFloat(s[240:253])
	r.ValorRecebido = parseFloat(s[253:266])
	r.JurosMora = parseFloat(s[266:279])
	r.OutrosRecebimentos = parseFloat(s[279:292])
	r.DataCredito = parseDate(s[295:301])
	seq, _ := strconv.Atoi(s[394:400])
	r.Sequencial = seq
	return
}

func (b SantanderCNAB400) nossoNumero(n string) string {
	nn := nossoNumero(n)
	return cobranca.Zeros(nn, 14)
}

func (b SantanderCNAB400) ocorrencia(co string) (string, string) {
	mapOcorrencia := map[string]string{
		"02": "entrada tít. confirmada",
		"03": "entrada tít. rejeitada",
		"06": "liquidação",
		"07": "liquidação por conta",
		"08": "liquidação por saldo",
		"09": "baixa automática",
		"10": "tít. baix. conf. instrução",
		"16": "tít. já baixado/liquidado",
		"17": "liquidado em cartório",
		"20": "",
		"28": "Manutenção de Título Vencido",
		"46": "",
		"72": "Alteração de Tipo"}

	desc := mapOcorrencia[co]
	switch co {
	case "02":
		return entrada, desc
	case "03":
		return recusa, desc
	case "05",
		"06",
		"07",
		"08",
		"16",
		"17",
		"46":
		return liquidado, desc
	case "09",
		"10",
		"20":
		return baixa, desc
	case "28":
		return manutencaoTituloVencido, desc
	case "72":
		return alteracaoTipo, desc
	}
	return "", ""
}
