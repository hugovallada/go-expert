package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hugovallada/go-expert/api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDb := NewProduct(db)

	products, err := productDb.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDb.FindAll(2, 10, "no_asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDb.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)

	products, err = productDb.FindAll(1, 10, "desc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 23", products[0].Name)
	assert.Equal(t, "Product 14", products[9].Name)

	products, err = productDb.FindAll(0, 0, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 23)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)
}

func TestFindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.NoError(t, err)
	product, err = productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestUpdateProduct_ShouldReturnAnErrorWhenTheProductDoesNotExists(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Update(product)
	assert.Error(t, err)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
	product, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
	assert.Empty(t, product)
}

func TestDeleteProduct_ShouldReturnErrorWhenTheIdDoeNotExissts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.Error(t, err)
}
