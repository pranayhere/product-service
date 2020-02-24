package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"product-service/product"
)

func main() {
	fmt.Println("First GoLang REST API")
	db := initDB()
	defer db.Close()

	productAPI := InitProductAPI(db)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindById)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})
	return db
}
