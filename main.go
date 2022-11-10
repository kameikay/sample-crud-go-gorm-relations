package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories"`
	// SerialNumber SerialNumber
	gorm.Model
}

// type SerialNumber struct {
// 	ID        int `gorm:"primaryKey"`
// 	Number    string
// 	ProductID int
// }

func main() {
	dsn := "root:root@tcp(localhost:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// CREATE CATEGORY
	// category := Category{
	// 	Name: "Cozinha",
	// }
	// db.Create(&category)

	// category2 := Category{
	// 	Name: "Eletronics",
	// }
	// db.Create(&category2)

	// // CREATE PRODUCT
	// db.Create(&Product{
	// 	Name:       "Cafeteira",
	// 	Price:      500.00,
	// 	Categories: []Category{category, category2},
	// })

	// CREATE SERIAL NUMBER
	// db.Create(&SerialNumber{
	// 	Number:    "1246437y34",
	// 	ProductID: 1,
	// })

	// var products []Product
	// // O Preload traz a categoria junto. Sem o preload, traz somente o produto.
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category, product.SerialNumber)
	// }

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ": ")
		for _, product := range category.Products {
			fmt.Println("-", product.Name)
		}
	}

}
