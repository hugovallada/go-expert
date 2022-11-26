package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.00)
	assert.Equal(t, 10.0, tax)
	assert.Nil(t, err)
}

func Test_CalculateTaxThrownError(t *testing.T) {
	tax, err := CalculateTax(0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
}
