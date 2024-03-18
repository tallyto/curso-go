package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primarykey"`
	Name string
	gorm.Model
}

type Product struct {
	ID         int `gorm:"primarykey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// create category

	// category := Category{Name: "Eletronics"}

	// db.Create(&category)

	// db.Create(&Product{
	// 	Name:       "Mouse",
	// 	Price:      2000,
	// 	CategoryID: 1,
	// })

	var products []Product

	db.Preload("Category").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}

}
