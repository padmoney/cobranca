package retorno

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/padmoney/cobranca"
)

const (
	CNAB240 = "cnab240"
	CNAB400 = "cnab400"

	alteracaoTipo           = "alteracao_tipo"
	baixa                   = "baixa"
	entrada                 = "entrada"
	liquidado               = "liquidado"
	manutencaoTituloVencido = "manutencao_titulo_vencido"
	recusa                  = "recusa"

	arquivoRetornoNaoEncontrado = "Arquivo de retorno não encontrado."
	layoutRetornoNaoSuportado   = "Layout do arquivo de retorno não é suportado."
)

var (
	layoutsValidos = map[string][]string{
		cobranca.CodigoBancoBrasil: []string{CNAB400},
		cobranca.CodigoSantander:   []string{CNAB400},
	}
)

type RetornoBanco interface {
	LerLinha(linha string) (Registro, error)
}

type Retorno struct {
	banco  string
	layout string
}

func New(banco, layout string) (Retorno, error) {
	r := Retorno{banco, layout}
	return r, r.AssertValid()
}

func (r Retorno) AssertValid() error {
	lv := layoutsValidos[r.banco]
	if lv == nil {
		return errors.New(cobranca.BancoNaoSuportado)
	}
	if !contains(r.layout, lv) {
		return errors.New(layoutRetornoNaoSuportado)
	}
	return nil
}

func (r Retorno) Read(path string) (registros []Registro, err error) {
	if !fileExists(path) {
		err = errors.New(arquivoRetornoNaoEncontrado)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linhas := strings.Split(scanner.Text(), "\n")
		r, e := r.read(linhas)
		if e != nil {
			err = e
			break
		}
		registros = append(registros, r...)
	}
	return
}

func (r Retorno) ReadFile(content []byte) ([]Registro, error) {
	linhas := strings.Split(string(content), "\n")
	return r.read(linhas)
}

func (r Retorno) read(linhas []string) (registros []Registro, err error) {
	var rb RetornoBanco
	switch r.banco {
	case cobranca.CodigoBancoBrasil:
		rb = NewBancoBrasilCNAB400()
	}

	for _, l := range linhas {
		registro, err := rb.LerLinha(l)
		if err != nil {
			break
		}
		if registro.ID == idRegistroDetalhe {
			registros = append(registros, registro)
		}
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
