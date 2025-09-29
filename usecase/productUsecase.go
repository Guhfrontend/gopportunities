package usecase

import (
	"fmt"

	"github.com/Guhfrontend/gopportunities/model"
)

type ProductUsecase struct{}

func NewProductUsecase() *ProductUsecase {
	return &ProductUsecase{}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	products := []model.Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 20.99},
		{ID: 3, Name: "Product 3", Price: 30.99},
	}
	return products, nil
	fmt.Println("Fetching products...")
}
