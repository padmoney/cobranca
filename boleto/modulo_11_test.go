package boleto

import (
	"testing"
)

func TestCalculateModulo11(t *testing.T) {
	m := map[string]string{
		"4556":              "1",
		"3680":              "3",
		"05009401448":       "1",
		"12387987777700168": "2",
		"85068014982":       "9",
		"0317847":           "1",
		"15399":             "0",
		"61900":             "0",
		"123456789012":      "3"}

	for v, expected := range m {
		got, _ := Modulo11{}.Calculate(v)
		if got != expected {
			t.Errorf("Expected '%s' got '%s' to '%s'", expected, got, v)
		}
	}
}

func TestCalculareModulo11WithMap(t *testing.T) {
	expected := "X"
	m := map[int]string{
		10: "X"}
	got, _ := Modulo11{
		Mapping: m}.Calculate("4556")
	if got != expected {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}

func TestMod11WithError(t *testing.T) {
	data := "X317847"
	_, err := Modulo11{}.Calculate(data)
	if err == nil {
		t.Errorf("Should not be possible to get a number representation of '%s'", data)
	}
}
