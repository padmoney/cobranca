package cobranca

type Beneficiario struct {
	Nome      string
	Documento string
}

func (b Beneficiario) TipoInscricao() string {
	return tipoInscricao(b.Documento)
}

func (b Beneficiario) GetDocumento() string {
	return SemMascara(b.Documento)
}
