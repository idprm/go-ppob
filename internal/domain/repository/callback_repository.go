package repository

import "gorm.io/gorm"

type CallbackRepository struct {
	db *gorm.DB
}

func NewCallbackRepository(
	db *gorm.DB,
) *CallbackRepository {
	return &CallbackRepository{
		db: db,
	}
}

type ICallbackRepository interface {
}
