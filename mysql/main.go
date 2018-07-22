package main

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/guobin8205/golearn/mysql/models"
)

func main(){
	db,err := gorm.Open("mysql","root:guobin@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Product{})

	// Create
	db.Create(&models.Product{Code: "L1212", Price: 1000})

	// Read
	var product models.Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}
