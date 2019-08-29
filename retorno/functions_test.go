package retorno

import "testing"

func TestNossoNumero(t *testing.T) {
	got := nossoNumero("1234567890")
	expected := "123456789-0"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
