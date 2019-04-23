package boleto

import (
	"strconv"
)

type Modulo11 struct {
	Mapping map[int]string
}

func (m Modulo11) Calculate(s string) (string, error) {
	data, err := Explode(s)
	if err != nil {
		return "", err
	}

	return m.computeCheckDigit(data), nil
}

func (m Modulo11) computeCheckDigit(data []int) string {
	var sum int

	for i, f := len(data)-1, 2; i >= 0; i, f = i-1, f+1 {
		sum += data[i] * f
		if f == 9 {
			f = 1
		}
	}

	mod := int(sum % 11)
	r := 0
	if mod != 0 {
		r = 11 - mod
	}
	if v := m.mapValue(r); v != "" {
		return v
	}
	if r > 9 {
		return "1"
	}
	return strconv.Itoa(r)
}

func (m Modulo11) mapValue(i int) string {
	if v := m.Mapping[i]; v != "" {
		return v
	}
	return ""
}
