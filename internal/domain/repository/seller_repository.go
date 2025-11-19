package repository

import "gorm.io/gorm"

type SellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(
	db *gorm.DB,
) *SellerRepository {
	return &SellerRepository{
		db: db,
	}
}

type ISellerRepository interface {
}
