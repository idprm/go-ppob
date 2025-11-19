package repository

import "gorm.io/gorm"

type BuyerRepository struct {
	db *gorm.DB
}

func NewBuyerRepository(
	db *gorm.DB,
) *BuyerRepository {
	return &BuyerRepository{
		db: db,
	}
}

type IBuyerRepository interface {
}
