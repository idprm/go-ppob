package repository

import "gorm.io/gorm"

type APIKeyRepository struct {
	db *gorm.DB
}

func NewAPIKeyRepository(db *gorm.DB) *APIKeyRepository {
	return &APIKeyRepository{
		db: db,
	}
}

type IAPIKeyRepository interface {
}
