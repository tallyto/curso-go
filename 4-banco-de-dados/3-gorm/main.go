package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primarykey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&Product{})

	// create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "Mouse", Price: 100},
	// 	{Name: "Keyboard", Price: 200},
	// 	{Name: "Monitor", Price: 300},
	// 	{Name: "Webcam", Price: 400},
	// 	{Name: "Microphone", Price: 500},
	// 	{Name: "Headset", Price: 600},
	// 	{Name: "Speaker", Price: 700},
	// 	{Name: "Pen Drive", Price: 800},
	// 	{Name: "Memory Card", Price: 900},
	// 	{Name: "USB Cable", Price: 1000},
	// 	{Name: "HDMI Cable", Price: 1100},
	// 	{Name: "DVD Drive", Price: 1200},
	// 	{Name: "SSD Drive", Price: 1300},
	// 	{Name: "Flash Drive", Price: 1400},
	// 	{Name: "Printer", Price: 1500},
	// 	{Name: "Scanner", Price: 1600},
	// }
	// db.Create(&products)

	// select one

	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Monitor")
	// fmt.Println(product)

	// select all

	var products []Product

	db.Find(&products)

	for _, product := range products {
		fmt.Println(product)
	}

}
