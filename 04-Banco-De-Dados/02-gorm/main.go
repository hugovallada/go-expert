package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	//SerialNumber SerialNumber
	gorm.Model
}

// type SerialNumber struct {
// 	ID        int `gorm:"primaryKey"`
// 	Number    string
// 	ProductID int
// }

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	tx := db.Begin() // Inicia a transaction
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletrodomestico"
	tx.Debug().Save(&c)
	tx.Commit()

	// // create category
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// category2 := Category{Name: "Cozinha"}
	// db.Create(&category2)

	// product := Product{
	// 	Name:       "Panela Eletrica",
	// 	Price:      100.00,
	// 	Categories: []Category{category, category2},
	// }
	// create product
	// db.Create(&product)

	// product2 := Product{
	// 	Name:       "Fritadeira Eletrica",
	// 	Price:      3000.00,
	// 	Categories: []Category{category, category2},
	// }

	//db.Create(&product2)

	// create a serialNumber
	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: int(product.ID),
	// })
	// // Belongs To
	// var products []Product
	// db.Preload("Category").Find(&products)

	// HasOne
	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)

	//var categories []Category

	// Problemas do hasMany
	// Pegar o relacionamento do products e serialNumber
	//err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error

	// Many To Many
	// err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	// if err != nil {
	// 	panic(err)
	// }

	// for _, cate := range categories {
	// 	fmt.Println(cate.Name, ":")
	// 	for _, pr := range cate.Products {
	// 		fmt.Println(pr.Name)
	// 	}
	// }

	// for _, product := range products {
	// 	fmt.Println(product.Name, product.CategoryID, product.Category.Name)
	// }

	// create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000,
	// })

	// // create batch
	// products := []Product{
	// 	{Name: "Roupa", Price: 50},
	// 	{Name: "Mouse", Price: 60},
	// 	{Name: "Coca-Cola", Price: 7.90},
	// }

	// db.Create(products)

	// // Select One
	// var p Product
	// //db.First(&p, 1)
	// db.First(&p, "name = ?", "Mouse")
	// fmt.Println(p)

	// // Select All
	// var multipleProducts []Product
	// db.Find(&multipleProducts)

	// for _, prd := range multipleProducts {
	// 	fmt.Println(prd)
	// }

	// // Limit
	// fmt.Println("Limit")
	// var limitedProducts []Product
	// db.Limit(2).Find(&limitedProducts)
	// fmt.Println(limitedProducts)

	// // Offset
	// fmt.Println("Offset")
	// var offsetProducts []Product
	// db.Limit(10).Offset(2).Find(&offsetProducts)
	// fmt.Println(offsetProducts)

	// // Where
	// fmt.Println("Where")
	// var pWhere []Product
	// db.Where("price > ?", 100).Find(&pWhere)
	// fmt.Println(pWhere)

	// // Alterando e Removendo produtos

	// var pd Product
	// db.First(&pd, 1)

	// pd.Price = 10000
	// pd.Name = "New Product"
	// db.Save(&pd)

	// var updatedProduct Product
	// db.First(&updatedProduct, 1)
	// fmt.Println(updatedProduct)

	// db.Delete(&updatedProduct)

	// Soft Delete

	//create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000,
	// })

	// var pd Product
	// db.First(&pd, 1)

	// pd.Price = 10000
	// pd.Name = "New Product"
	// db.Save(&pd)

	// var pd Product
	// db.First(&pd, 1)
	// db.Delete(&pd)
}

/**
* Lock Optimista vs Pessimista
*
* Otimista, verifica a versão dos dados, se tiver a msm commita, se não, faz rollback e começa a alteração com os novos dados
* Pessimitas, Locka a tabela, a linha do banco e ngm consegue alterar
*
* Também é um exemplo de concorrência
 */
