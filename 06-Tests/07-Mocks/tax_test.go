package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("Save", 10.0).Return(nil)
	//repository.On("Save", mock.Anything).Return(nil)
	//repository.On("Save", mock.Anything).Return(nil).Times(1) define qnts vezes pode ser chamado
	repository.On("Save", 0.0).Return(errors.New("error saving tax"))
	err := CalculateTaxAndSave(1000.00, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")
	repository.AssertExpectations(t)
	//repository.AssertNumberOfCalls(t, "Save", 3) Spy.. garante q o método do repositório foi chamado 3 vezes
}

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
