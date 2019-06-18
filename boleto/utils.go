package boleto

import (
	"math"
	"strconv"
	"time"
)

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func Explode(str string) ([]int, error) {
	var data []int
	for _, s := range str {
		n, err := strconv.Atoi(string(s))
		if err != nil {
			return data, err
		}
		data = append(data, n)
	}
	return data, nil
}

func ParseDate(s string) time.Time {
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

func ParseFloat(s string) float64 {
	valor, _ := strconv.ParseFloat(s, 64)
	return valor / 100
}

func Round(v float64, places int) float64 {
	const roundOn = .5

	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * v
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}
