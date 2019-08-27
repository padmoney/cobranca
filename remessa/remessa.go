package remessa

import (
	"errors"
	"fmt"
	"strings"

	"github.com/padmoney/cobranca"
)

const (
	CNAB400 = "cnab400"

	ComandoCancelamento = "cancelamento"
	ComandoRegistro     = "registro"

	layoutRemessaNaoSuportado = "Layout do arquivo de remessa não é suportado."
)

type geradorRemessa interface {
	Trailler() string
	Header() string
	Pagamentos(pagamentos []Pagamento) ([]string, error)
}

type Remessa struct {
	layout     string
	conta      cobranca.Conta
	params     Params
	sequencial int64
	gerador    geradorRemessa

	pagamentos []Pagamento
}

func New(layout string, conta cobranca.Conta, params Params, sequencial int64) (*Remessa, error) {
	remessa := &Remessa{
		conta:      conta,
		params:     params,
		sequencial: sequencial}
	if strings.ToLower(layout) != CNAB400 {
		return remessa, errors.New(layoutRemessaNaoSuportado)
	}
	switch conta.Banco {
	case cobranca.CodigoBancoBrasil:
		remessa.gerador = NewBancoBrasilCNAB400(conta, params)
	case cobranca.CodigoSantander:
		remessa.gerador = NewSantanderCNAB400(conta, params)
	default:
		return remessa, errors.New(cobranca.BancoNaoSuportado)
	}
	return remessa, nil
}

func (r *Remessa) AddPagamento(p Pagamento) error {
	r.pagamentos = append(r.pagamentos, p)
	return nil
}

func (r Remessa) NomeArquivo() string {
	return fmt.Sprintf("REMESSA_%d.rem", r.sequencial)
}

func (r Remessa) Linhas() (linhas []string, err error) {
	header := r.gerador.Header()
	linhas = append(linhas, header)

	pagamentos, err := r.gerador.Pagamentos(r.pagamentos)
	if err != nil {
		return
	}
	for _, p := range pagamentos {
		linhas = append(linhas, cobranca.Sanitize(p))
	}

	trailler := r.gerador.Trailler()
	linhas = append(linhas, trailler)
	return
}
