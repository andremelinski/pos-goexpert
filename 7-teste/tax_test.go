package tax_test

import (
	"testing"

	tax "github.com/andremelinski/pos-goexpert/7-test"
)

func TestCalculateTaxt(t *testing.T) {
	amount := 500.0
	expectResult := 5.0

	result := tax.CalculateTaxt(amount)

	if result != expectResult {
		t.Errorf("Excpected %f but got %f", expectResult, result)
	}
}

