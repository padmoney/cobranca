package boleto

import (
	"reflect"
	"testing"
	"time"
)

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
