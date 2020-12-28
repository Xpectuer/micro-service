package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{ID: 1,
		Name:  "alex",
		Price: 1.00,
		SKU:   "abc-abc-def",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
