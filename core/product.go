package core

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string
	Quantity int64
}

type FindOneProductResponse struct {
	ProductName string `json:"product_name"`
	Quantity    int64  `json:"quantity"`
}
