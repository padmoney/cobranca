package remessa

import (
	"math"
	"strconv"
	"time"

	"github.com/padmoney/cobranca"
)

func DataFormatada(data time.Time) string {
	return data.Format("020106")
}

func ValorFormatado(valor float64, q, d int) string {
	p := math.Pow(10, float64(d))
	valor = math.Round(valor * p)
	s := strconv.FormatFloat(valor, 'f', -1, 64)
	return cobranca.Zeros(s, q+d)
}
