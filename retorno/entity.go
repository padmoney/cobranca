package retorno

import "time"

const (
	idRegistroHeader  = "0"
	idRegistroDetalhe = "7"
	idRegistroTrailer = "9"
)

type Registro struct {
	ID                 string
	TipoCobranca       string
	DataOcorrencia     time.Time
	Valor              float64
	DataVencimento     time.Time
	Numero             string
	NossoNumero        string
	ValorTarifa        float64
	Ocorrencia         string
	TipoOcorrencia     string
	DataLiquidacao     time.Time
	EspecieTitulo      string
	OutrasDespesas     float64
	JurosAtraso        float64
	TaxaIOF            float64
	ValorAbatimento    float64
	DescontoConcedido  float64
	ValorRecebido      float64
	JurosMora          float64
	OutrosRecebimentos float64
	DataCredito        time.Time
	Sequencial         int
}
