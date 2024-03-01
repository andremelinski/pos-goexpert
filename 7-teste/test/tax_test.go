package tax_test

import (
	"testing"

	tax "github.com/andremelinski/pos-goexpert/7-test/test"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectResult := 5.0

	result := tax.CalculateTax(amount)

	if result != expectResult {
		t.Errorf("Excpected %f but got %f", expectResult, result)
	}
}


func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}
	table := []calcTax{
		{500.0, 5.0},
		{100.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := tax.CalculateTax(item.amount)
	
		if result != item.expect {
			t.Errorf("Excpected %f but got %f", item.expect, result)
		}
	}
}


func BenchTestCalculateTax(b *testing.B) {
	type calcTax struct {
		amount, expect float64
	}
	for i := 0; i < b.N; i++ {
		tax.CalculateTax(500.0)
	}
}

func BenchTestCalculateTax2(b *testing.B) {
	type calcTax struct {
		amount, expect float64
	}
	for i := 0; i < b.N; i++ {
		tax.CalculateTax(500.0)
	}
}


func FuzzCalculateTax(f *testing.F){
	// mock data
	seed := []float64{-1,-5,-2.5,500.0,1000.0,151.78}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func (t *testing.T, amount float64)  {
		result := tax.CalculateTax(amount)
		if amount <=0 && result !=0{
			t.Errorf("Received %f but expected 0", result)
		}
		if amount > 20000 && result !=20{
			t.Errorf("Received %f but expected 20", result)
		}
	})
}