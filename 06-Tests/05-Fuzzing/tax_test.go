package tax

import "testing"

// FUZZING = TESTE DE MUTAÇÃO

func Fuzz_CalculateTax(f *testing.F) {
	// Seeding - Dar informações
	seed := []float64{-1.0, -2.0, -2.5, -500.0, 1000.0, 1500.0, 1501.0}

	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %.2f, but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Received %.2f, but expected 20", result)
		}
	})
}

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
		{0.0, 0.0},
	}

	for _, entry := range table {
		result := CalculateTax(entry.ammount)
		if result != entry.expected {
			t.Errorf("Expected %.2f, but got %.2f", entry.expected, result)
		}
	}
}

func Benchmark_CalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func Benchmark_CalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}
