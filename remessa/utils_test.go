package remessa

import (
	"testing"
	"time"
)

func TestDataFormatada(t *testing.T) {
	layout := "2006-01-02"
	str := "1984-11-09"
	date, err := time.Parse(layout, str)
	if err != nil {
		t.Errorf("There should not be an error, error: %s", err)
	}

	expected := "091184"
	if got := DataFormatada(date); got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
