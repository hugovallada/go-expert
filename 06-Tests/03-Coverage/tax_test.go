package tax

import "testing"

func Test_CalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}

func Test_CalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		ammount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{20, 5.0},
	}

	for _, entry := range table {
		result := CalculateTax(entry.ammount)
		if result != entry.expected {
			t.Errorf("Expected %.2f, but got %.2f", entry.expected, result)
		}
	}
}
