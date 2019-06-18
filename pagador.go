package cobranca

type Avalista struct {
	Nome      string
	Documento string
}

type Pagador struct {
	Nome      string
	Documento string
	Endereco  string
	Bairro    string
	Cidade    string
	UF        string
	CEP       string
}

func (p Pagador) GetDocumento() string {
	return OnlyNumbers(p.Documento)
}

func (p Pagador) PessoaFisica() bool {
	doc := p.GetDocumento()
	return len(doc) == 11
}

func (p Pagador) TipoInscricao() string {
	return tipoInscricao(p.Documento)
}

// TipoInscricao retorna o tipo de inscrição
// 00 - ISENTO
// 01 - CPF
// 02 - CNPJ
func tipoInscricao(doc string) string {
	doc = SemMascara(doc)
	switch len(doc) {
	case 11:
		return "01"
	case 14:
		return "02"
	default:
		return "00"
	}
}
