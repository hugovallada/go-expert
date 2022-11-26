package tax

import "testing"

func Test_CalculateTax(t *testing.T) {
	amount := 500.0
	expected := 6.0

	result := CalculateTax(amount)
	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}
