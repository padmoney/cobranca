package boleto

import (
	"strconv"
)

type Modulo10 struct {
}

func (m Modulo10) Calculate(s string) (string, error) {
	data, err := Explode(s)
	if err != nil {
		return "", err
	}

	return m.computeCheckDigit(data), nil
}

func (m Modulo10) computeCheckDigit(data []int) string {
	var sum int
	multiplier := 2
	reversed := reverse(data)

	for _, i := range reversed {
		c := i * multiplier
		s := sumDigits(c)
		sum = sum + s
		if multiplier == 2 {
			multiplier = 1
		} else {
			multiplier = 2
		}
	}
	r := 10 - (sum % 10)
	if r == 10 {
		return "0"
	}
	return strconv.Itoa(r)
}

func reverse(n []int) []int {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return n
}

func sumDigits(d int) int {
	var sum int
	s := strconv.Itoa(d)
	data, _ := Explode(s)
	for _, i := range data {
		sum = sum + i
	}
	return sum
}
