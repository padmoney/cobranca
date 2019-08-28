package retorno

import (
	"errors"
	"os"

	"github.com/padmoney/cobranca"
)

const (
	CNAB240 = "cnab240"
	CNAB400 = "cnab400"

	arquivoRetornoNaoEncontrado = "Arquivo de retorno não encontrado."
	layoutRetornoNaoSuportado   = "Layout do arquivo de retorno não é suportado."
)

var (
	bancosValidos = []string{
		cobranca.CodigoBancoBrasil,
		cobranca.CodigoSantander,
	}
	layoutsValidos = map[string][]string{
		cobranca.CodigoBancoBrasil: []string{CNAB400},
		cobranca.CodigoSantander:   []string{CNAB400},
	}
)

type Retorno struct {
	banco  string
	layout string
}

func New(banco, layout string) (Retorno, error) {
	r := Retorno{banco, layout}
	return r, r.AssertValid()
}

func (r Retorno) AssertValid() error {
	if !contains(r.banco, bancosValidos) {
		return errors.New(cobranca.BancoNaoSuportado)
	}
	lv := layoutsValidos[r.banco]
	if !contains(r.layout, lv) {
		return errors.New(layoutRetornoNaoSuportado)
	}
	return nil
}

func (r Retorno) Read(path string) (items []Item, err error) {
	if !fileExists(path) {
		err = errors.New(arquivoRetornoNaoEncontrado)
		return
	}
	return
}

func contains(v string, a []string) bool {
	for _, i := range a {
		if v == i {
			return true
		}
	}
	return false
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
