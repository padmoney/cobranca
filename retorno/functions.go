package retorno

import (
	"fmt"
	"strconv"
	"time"
)

func nossoNumero(s string) string {
	l := len(s) - 1
	return fmt.Sprintf("%s-%s", s[:l], s[l:])
}

func parseDate(s string) time.Time {
	if s == "000000" {
		return time.Time{}
	}
	if len(s) != 6 {
		return time.Time{}
	}

	f := "020106"
	t, _ := time.Parse(f, s)
	return t
}

func parseFloat(s string) float64 {
	valor, _ := strconv.ParseFloat(s, 64)
	return valor / 100
}
