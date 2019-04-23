package boleto

import (
	"reflect"
	"testing"
	"time"
)

func TestBrancos(t *testing.T) {
	qote := "a long time ago"
	got := Brancos(qote, 20)
	expected := "A LONG TIME AGO     "
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}

	got = Brancos("a long time ago", 8)
	expected = "TIME AGO"
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}

func TestDate(t *testing.T) {
	now := time.Now()
	expected := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	got := Date(now.Year(), int(now.Month()), now.Day())
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected '%v' got '%v'", expected, got)
	}
}

func TestExplode(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	got, err := Explode("1234567890")

	if err != nil {
		t.Errorf("Should be nil")
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected '%v' got '%v'", expected, got)
	}
}

func TestOnlyNumbers(t *testing.T) {
	expected := "1234567890"
	got := OnlyNumbers("1a2!3@4#5%6&7*8()9-0{}ç^;.,?/<>|")
	if got != expected {
		t.Errorf("Expected '%v' got '%v'", expected, got)
	}
}

func TestParseDate(t *testing.T) {
	expected := time.Time{}
	got := ParseDate("000000")
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected.Format("2006-01-02"), got.Format("2006-01-02"))
	}

	got = ParseDate("12345")
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected.Format("2006-01-02"), got.Format("2006-01-02"))
	}

	expected = time.Date(1984, time.Month(9), 11, 0, 0, 0, 0, time.UTC)
	got = ParseDate("110984")
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected.Format("2006-01-02"), got.Format("2006-01-02"))
	}
}

func TestParseFloat(t *testing.T) {
	expected := 2027.50
	got := ParseFloat("00000000202750")
	if expected != got {
		t.Errorf("Expected '%f' got '%f'", expected, got)
	}
}

func TestRound(t *testing.T) {
	m := map[float64]float64{
		24.8:       24.80,
		123.555555: 123.56,
		123.558:    123.56,
		69.99:      69.99,
		69.98:      69.98,
	}

	for v, expected := range m {
		got := Round(v, 2)
		if got != expected {
			t.Errorf("Expected '%f' got '%f'", expected, got)
		}
	}
}

func TestSanitize(t *testing.T) {
	equivalents := map[string]string{
		"À":         "A",
		"Á":         "A",
		"Â":         "A",
		"Ã":         "A",
		"Ä":         "A",
		"Å":         "AA",
		"Æ":         "AE",
		"Ç":         "C",
		"È":         "E",
		"É":         "E",
		"Ê":         "E",
		"Ë":         "E",
		"Ì":         "I",
		"Í":         "I",
		"Î":         "I",
		"Ï":         "I",
		"Ð":         "D",
		"Ł":         "L",
		"Ñ":         "N",
		"Ò":         "O",
		"Ó":         "O",
		"Ô":         "O",
		"Õ":         "O",
		"Ö":         "O",
		"Ø":         "OE",
		"Ù":         "U",
		"Ú":         "U",
		"Ü":         "U",
		"Û":         "U",
		"Ý":         "Y",
		"Þ":         "TH",
		"ß":         "SS",
		"à":         "A",
		"á":         "A",
		"â":         "A",
		"ã":         "A",
		"ä":         "A",
		"å":         "AA",
		"æ":         "AE",
		"ç":         "C",
		"è":         "E",
		"é":         "E",
		"ê":         "E",
		"ë":         "E",
		"ì":         "I",
		"í":         "I",
		"î":         "I",
		"ï":         "I",
		"ð":         "D",
		"ł":         "L",
		"ñ":         "N",
		"ń":         "N",
		"ò":         "O",
		"ó":         "O",
		"ô":         "O",
		"õ":         "O",
		"ō":         "O",
		"ö":         "O",
		"ø":         "OE",
		"ś":         "S",
		"ù":         "U",
		"ú":         "U",
		"û":         "U",
		"ū":         "U",
		"ü":         "U",
		"ý":         "Y",
		"þ":         "TH",
		"ÿ":         "Y",
		"ż":         "Z",
		"Œ":         "OE",
		"œ":         "OE",
		"żż":        "ZZ",
		"Hello, 世界": "HELLO, ??",
	}

	for k, expected := range equivalents {
		got := Sanitize(k)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestStrPad(t *testing.T) {
	expected := "0000012345-6"
	got := StrPad("0000012345-6", 12, "0", StrPadLeft)
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}

	expected = "0000012345-6"
	got = StrPad("12345-6", 12, "0", StrPadLeft)
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}

	expected = "6789012345-6"
	got = StrPad("0123456789012345-6", 12, "0", StrPadLeft)
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}

	expected = "Padmoney  "
	got = StrPad("Padmoney", 10, " ", StrPadRight)
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}

	expected = "Clap Clap Clap Clap "
	got = StrPad("", 20, "Clap ", StrPadRight)
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}

	expected = "Padmoney"
	got = StrPad("Padmoney Payment Security", 8, " ", StrPadRight)
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}

func TestZeros(t *testing.T) {
	got := Zeros("12345-6", 12)
	expected := "0000012345-6"
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}

	got = Zeros("1234567890", 5)
	expected = "67890"
	if expected != got {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}
