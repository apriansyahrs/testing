package repository

import (
	"gin-railway/core"
	"gin-railway/database"
)

type ProductRepo struct {
	postgres *database.Postgres
}

func NewProductRepo(db *database.Postgres) *ProductRepo {
	return &ProductRepo{
		postgres: db,
	}
}

func (r *ProductRepo) FindProductByID(id uint) (*core.Product, error) {
	product := core.Product{}
	err := r.postgres.DB.First(&product, "id=?", id).Error

	return &product, err
}
