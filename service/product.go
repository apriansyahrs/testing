package service

import (
	"gin-railway/core"
	"gin-railway/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepo
}

func NewProductService(productRepo *repository.ProductRepo) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) FindOneProduct(id int) (*core.FindOneProductResponse, error) {
	product, err := s.productRepo.FindProductByID(uint(id))
	if err != nil {
		return nil, err
	}

	resp := &core.FindOneProductResponse{
		ProductName: product.Name,
		Quantity:    product.Quantity,
	}

	return resp, nil
}
