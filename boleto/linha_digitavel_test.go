package boleto

import (
	"testing"
)

func TestLinhaDigitavel(t *testing.T) {
	m := map[string]string{
		"00195709300000545330000002962087000000000217": "00190000090296208700900000002170570930000054533",
		"00193000000000005450000002962087000000000217": "00190000090296208700900000002170300000000000545",
		"00197000000000005450000002962087000000000117": "00190000090296208700900000001172700000000000545",
		"00199000000000042090000002962087000000000217": "00190000090296208700900000002170900000000004209",
		"00198000000000176720000002962087000000000317": "00190000090296208700900000003178800000000017672",
		"00199000000001540360000002962087000000000417": "00190000090296208700900000004176900000000154036",
		"00196000000017852930000002962087000000000517": "00190000090296208700900000005173600000001785293",
		"03394740000012345679897033512345678901230101": "03399897093351234567089012301019474000001234567"}

	for cb, ld := range m {
		got, _ := LinhaDigitavel{}.Calculate(cb)
		if got != ld {
			t.Errorf("Expected '%s' got '%s'", ld, got)
		}
	}
}

func TestErrorCampo1(t *testing.T) {
	got, err := LinhaDigitavel{}.Calculate("A0195709300000545330000002962087000000000217")
	if err == nil {
		t.Errorf("Should'n be nil")
	}
	if got != "" {
		t.Errorf("Should'n be empty")
	}
}

func TestErrorCampo2(t *testing.T) {
	got, err := LinhaDigitavel{}.Calculate("001930000000000054500000_2962087000000000217")
	if err == nil {
		t.Errorf("Should'n be nil")
	}
	if got != "" {
		t.Errorf("Should'n be empty")
	}
}

func TestErrorCampo3(t *testing.T) {
	got, err := LinhaDigitavel{}.Calculate("0019700000000000545000000296208700+000000117")
	if err == nil {
		t.Errorf("Should'n be nil")
	}
	if got != "" {
		t.Errorf("Should'n be empty")
	}
}

func TestErrorCampo4(t *testing.T) {
	got, err := LinhaDigitavel{}.Calculate("0019-000000000176720000002962087000000000317+000000117")
	if err == nil {
		t.Errorf("Should'n be nil")
	}
	if got != "" {
		t.Errorf("Should'n be empty")
	}
}

func TestErrorCampo5(t *testing.T) {
	got, err := LinhaDigitavel{}.Calculate("00196/00000017852930000002962087000000000517")
	if err == nil {
		t.Errorf("Should'n be nil")
	}
	if got != "" {
		t.Errorf("Should'n be empty")
	}
}
