package boleto

import (
	"testing"
)

func TestNossoNumeroBancoBrasil(t *testing.T) {
	values := map[string]string{
		"4556":              "X",
		"3680":              "3",
		"29948060000035051": "2",
	}
	for k, v := range values {
		got, err := NossoNumero(k).BancoBrasil().DV()
		if err != nil {
			t.Errorf("There should not be an error, error: %s", err)
		}
		if got != v {
			t.Errorf("Expected '%s', got '%s'", v, got)
		}

	}
}
