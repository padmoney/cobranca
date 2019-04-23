package boleto

import (
	"reflect"
	"testing"
)

func TestCalculateModulo10(t *testing.T) {
	m := map[string]string{
		"0":          "0",
		"001905009":  "5",
		"4014481606": "9",
		"0680935031": "4"}

	for v, expected := range m {
		got, _ := Modulo10{}.Calculate(v)
		if got != expected {
			t.Errorf("Expected '%s' got '%s' to '%s'", expected, got, v)
		}
	}
}

func TestCalculateModulo10WithError(t *testing.T) {
	data := "X317847"
	_, err := Modulo10{}.Calculate(data)
	if err == nil {
		t.Errorf("Should not be possible to get a number representation of '%s'", data)
	}
}

func TestReverse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	expected := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	got := reverse(a)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected '%v' got '%v'", expected, got)
	}
}
