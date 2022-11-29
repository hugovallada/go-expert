package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
	assert.NotNil(t, p.CreatedAt)
}

type TestTable struct {
	TestDescription string
	EntryValue      Product
	ExpectedError   error
}

func Test_ProductValidate_WithInvalidData(t *testing.T) {
	testSuite := []TestTable{
		{"should return an error when the name isn't passed", Product{Name: "", Price: 10}, ErrNameIsRequired},
		{"should return an error when the price isn't passed", Product{Name: "Product", Price: 0}, ErrPriceIsRequired},
		{"should return an error when the price is invalid", Product{Name: "Product", Price: -20}, ErrInvalidPrice},
	}

	for _, suite := range testSuite {
		p, err := NewProduct(suite.EntryValue.Name, suite.EntryValue.Price)
		assert.Nil(t, p)
		assert.Equal(t, suite.ExpectedError, err)
	}
}
