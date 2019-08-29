package retorno

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/padmoney/cobranca/boleto"
)

type BancoBrasilCNAB400 struct {
}

func NewBancoBrasilCNAB400() BancoBrasilCNAB400 {
	return BancoBrasilCNAB400{}
}

func (b BancoBrasilCNAB400) LerLinha(linha string) (r Registro, err error) {
	if linha == "" {
		return
	}
	r.ID = linha[0:1]
	switch r.ID {
	case idRegistroHeader:
		if linha[1:9] != "2RETORNO" {
			err = errors.New("Tipo de Operação inválido.")
			return
		}
		if linha[9:19] != "01COBRANCA" {
			err = errors.New("Tipo de Serviço inválido.")
			return
		}
	case idRegistroDetalhe:
		r, err = b.lerDetalhe(r, linha)
	case idRegistroTrailer:
	}
	return
}

func (b BancoBrasilCNAB400) lerDetalhe(reg Registro, s string) (r Registro, err error) {
	r = reg

	r.Numero = strings.Replace(s[38:63], " ", "", -1)
	r.NossoNumero = b.nossoNumero(s[63:80])
	r.TipoCobranca = s[80:81]
	comando := s[108:110]
	r.TipoOcorrencia, r.Ocorrencia = b.ocorrencia(comando)
	if r.TipoOcorrencia == liquidado {
		r.DataLiquidacao = parseDate(s[110:116])
	}
	r.DataVencimento = parseDate(s[146:152])
	r.Valor = parseFloat(s[152:165])
	r.DataCredito = parseDate(s[175:181])
	r.ValorTarifa = parseFloat(s[181:188])
	r.OutrasDespesas = parseFloat(s[188:201])
	r.ValorAbatimento = parseFloat(s[227:240])
	r.DescontoConcedido = parseFloat(s[240:253])
	r.ValorRecebido = parseFloat(s[253:266])
	r.JurosMora = parseFloat(s[266:279])
	r.OutrosRecebimentos = parseFloat(s[279:292])
	seq, _ := strconv.Atoi(s[394:400])
	r.Sequencial = seq
	return
}

func (b BancoBrasilCNAB400) nossoNumero(n string) string {
	dv, _ := boleto.NossoNumero(n).BancoBrasil().DV()
	return fmt.Sprintf("%s-%s", n, dv)
}

func (b BancoBrasilCNAB400) ocorrencia(comando string) (tipo string, desc string) {
	m := map[string]string{
		"02": "Confirmação de Entrada de Boleto",
		"03": "Comando recusado",
		"05": "Liquidado sem registro (carteira 17-tipo4)",
		"06": "Liquidação Normal",
		"07": "Liquidação por Conta/Parcial",
		"08": "Liquidação por Saldo",
		"09": "Baixa de Titulo",
		"10": "Baixa Solicitada",
		"11": "Boletos em Ser",
		"12": "Abatimento Concedido",
		"13": "Abatimento Cancelado",
		"14": "Alteração de Vencimento do boleto",
		"15": "Liquidação em Cartório",
		"16": "Confirmação de alteração de juros de mora",
		"19": "Confirmação de recebimento de instruções para protesto",
		"20": "Débito em Conta",
		"21": "Alteração do Nome do Sacado",
		"22": "Alteração do Endereço do Sacado",
		"23": "Indicação de encaminhamento a cartório",
		"24": "Sustar Protesto",
		"25": "Dispensar Juros de mora",
		"26": "Alteração do número do boleto",
		"28": "Manutenção de titulo vencido",
		"31": "Conceder desconto",
		"32": "Não conceder desconto",
		"33": "Retificar desconto",
		"34": "Alterar data para desconto",
		"35": "Cobrar Multa",
		"36": "Dispensar Multa",
		"37": "Dispensar Indexador",
		"38": "Dispensar prazo limite para recebimento",
		"39": "Alterar prazo limite para recebimento",
		"41": "Alteração do número do controle do participante",
		"42": "Alteração do número do documento do sacado",
		"44": "Boleto pago com cheque devolvido",
		"46": "Boleto pago com cheque, aguardando compensação",
		"72": "Alteração de tipo de cobrança",
		"73": "Confirmação de Instrução de Parâmetro de Pagamento Parcial",
		"85": "Inclusão de Negativação",
		"86": "Exclusão de Negativação",
		"96": "Despesas de Protesto",
		"97": "Despesas de Sustação de Protesto",
		"98": "Débito de Custas Antecipadas",
	}
	desc = m[comando]
	switch comando {
	case "02":
		tipo = entrada
	case "03":
		tipo = recusa
	case "05", "06", "07", "08":
		tipo = liquidado
	case "09", "10":
		tipo = baixa
	case "28":
		tipo = recusa
	case "72":
		tipo = alteracaoTipo
	}
	return
}
