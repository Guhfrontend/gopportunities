package controller

import (
	"net/http"

	"github.com/Guhfrontend/gopportunities/model"
	"github.com/Guhfrontend/gopportunities/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) *ProductController {
	return &ProductController{
		ProductUsecase: usecase,
	}
}

func (pc *ProductController) GetProducts(c *gin.Context) {

	products := []model.Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 20.99},
		{ID: 3, Name: "Product 3", Price: 30.99},
	}
	c.JSON(http.StatusOK, products)
}
