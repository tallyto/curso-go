package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primarykey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

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

	// var products []Product

	// db.Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where

	// var products []Product
	// db.Limit(2).Find(&products)
	// db.Limit(2).Offset(2).Find(&products)
	// db.Where("price > ?", 1000).Find(&products)
	// db.Where("name LIKE ?", "%k%").Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update

	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2)

	// delete

	db.Delete(&p)
}
