package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Products []Product
	gorm.Model
}

type Product struct {
	ID           int `gorm:"primarykey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primarykey"`
	Number    string
	ProductID int
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// create category
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Mouse",
		Price:      2000,
		CategoryID: category.ID,
	})

	// create serial number

	// db.Create(&SerialNumber{
	// 	Number:    "123",
	// 	ProductID: 1,
	// })

	// var products []Product

	// db.Preload("Category").Preload("SerialNumber").Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println(" -", product.Name)
		}
	}

}
