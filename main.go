package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//products := []Product{
	//	{Name: "Notebook", Price: 2500.00},
	//	{Name: "Tablet", Price: 1200.00},
	//	{Name: "TV LG Smart", Price: 3400.00},
	//}
	//db.Create(&products)

	//var product Product
	//db.First(&product, 2)
	//fmt.Println(product)
	//var products []Product
	//db.Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	var products []Product
	db.Where("name LIKE ?", "%k%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
