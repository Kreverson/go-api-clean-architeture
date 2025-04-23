package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func NewProductController() ProductController {
	return ProductController{}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
		{ID: 3, Name: "Product 3", Price: 30.0},
	}
	ctx.JSON(http.StatusOK, products)
}
