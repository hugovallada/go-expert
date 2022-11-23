package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}

func main() {
	// Abrindo conexão
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert") // usuario:senha@tipo_conexao(host:porta)/nome_database
	if err != nil {
		panic(err)
	}
	defer db.Close() // fecha a conexão com o banco ao final
	product := NewProduct("Notebook", 1899.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 1500.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	returnedProduct, err := selectOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f\n", returnedProduct.Name, returnedProduct.Price)
	fmt.Println("BUSCANDO TUDO...")
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, prod := range products {
		fmt.Printf("Produto: %v, possui o preço de %.2f\n", prod.Name, prod.Price)

	}
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products (id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select * from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select * from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
