package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//camada de repositório
	ProductRepository := repository.NewProductRepository(dbConnection)

	//camada de usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	//camada de controller
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
