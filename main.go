package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&Product{})

	//Create
	db.Create(&Product{Code: "D42", Price: 100})

	//Read
	var product Product
	fmt.Println(db.First(&product, 1))                 //find product with integer primary key
	fmt.Println(db.First(&product, "code = ?", "D42")) //find product with code D42

	//Update - update product price to 200
	db.Model(&product).Update("Price", 200)

	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "C90"})

	//Delete Product
	db.Delete(&product, 1)

}
