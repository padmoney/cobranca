package remessa

import (
	"testing"
)

func TestParams(t *testing.T) {
	name := "Breno"
	company := "Padmoney"
	p := NewParams()
	p.Add("name", name)
	p.Add("company", company)
	if got := p.Get("name"); got != name {
		t.Errorf("Expected '%s', got '%s'", name, got)
	}
	if got := p.Get("company"); got != company {
		t.Errorf("Expected '%s', got '%s'", company, got)
	}

	// updating
	newName := "Bernardo"
	p.Add("name", newName)
	if got := p.Get("name"); got != newName {
		t.Errorf("Expected '%s', got '%s'", newName, got)
	}
}
