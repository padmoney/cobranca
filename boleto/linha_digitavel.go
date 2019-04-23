package boleto

import (
	"strconv"
)

type LinhaDigitavel struct {
	codigoBarras string
}

func (ld LinhaDigitavel) Calculate(cb string) (string, error) {
	ld.codigoBarras = cb

	campo1, err := ld.campo1()
	if err != nil {
		return "", err
	}

	campo2, err := ld.campo2()
	if err != nil {
		return "", err
	}

	campo3, err := ld.campo3()
	if err != nil {
		return "", err
	}

	campo4, err := ld.campo4()
	if err != nil {
		return "", err
	}
	campo5, err := ld.campo5()
	if err != nil {
		return "", err
	}

	linha := campo1 + campo2 + campo3 + campo4 + campo5
	return linha, nil
}

// Campo 1 da linha digitável
// a) Campo 1: AAABC.CCCCX
// A = Número Código da IF Destinatária no SILOC
// B = Código da moeda (9) -Real
// C = Posições 20 a 24 do código de barras
// X = DV do Campo 1 (calculado de acordo com o Módulo 10 – anexo V)
func (ld LinhaDigitavel) campo1() (string, error) {
	cd := ld.codigoBarras
	campo1 := cd[0:3] + cd[3:4] + cd[19:24]
	dv, err := Modulo10{}.Calculate(campo1)
	if err != nil {
		return "", err
	}
	return campo1 + dv, nil
}

// Campo 2 da linha digitável
// b) Campo 2: DDDDD.DDDDDY
// D = Posições 25 a 34 do código de barras
// Y = DV do Campo 2 (calculado de acordo com o Módulo 10 – anexo V)
func (ld LinhaDigitavel) campo2() (string, error) {
	cd := ld.codigoBarras
	campo2 := cd[24:34]
	dv, err := Modulo10{}.Calculate(campo2)
	if err != nil {
		return "", err
	}
	return campo2 + dv, nil
}

// Campo 3 da linha digitável
// c) Campo 3: EEEEE.EEEEEZ
// E = Posições 35 a 44 do código de barras
// Z = DV do Campo 3 (calculado de acordo com o Módulo 10 – anexo V)
func (ld LinhaDigitavel) campo3() (string, error) {
	cd := ld.codigoBarras
	campo3 := cd[34:44]
	dv, err := Modulo10{}.Calculate(campo3)
	if err != nil {
		return "", err
	}
	return campo3 + dv, nil
}

// Campo 4 da linha digitável
// d) Campo 4: K
// K = DV do código de barras (calculado de acordo com o Módulo 11 – anexo VI)
func (ld LinhaDigitavel) campo4() (string, error) {
	cd := ld.codigoBarras
	campo4 := cd[4:5]
	_, err := strconv.Atoi(string(campo4))
	if err != nil {
		return "", err
	}
	return campo4, nil
}

// Campo 5 da linha digitável
// e) Campo 5: UUUUVVVVVVVVVV
// U = Fator de Vencimento (cálculo conforme anexo IV)
// V = Valor do boleto de pagamento (com duas casas decimais, sem ponto e vírgula. Em caso de moeda variável, informar zeros)
func (ld LinhaDigitavel) campo5() (string, error) {
	cd := ld.codigoBarras
	campo5 := cd[5:9] + cd[9:19]
	_, err := strconv.Atoi(string(campo5))
	if err != nil {
		return "", err
	}
	return campo5, nil
}
