package boleto

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
