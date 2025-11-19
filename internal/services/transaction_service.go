package services

import (
	"github.com/idprm/go-ppob/internal/domain/repository"
	"gorm.io/gorm"
)

type TransactionService struct {
	transactionRepo repository.ITransactionRepository
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{
		transactionRepo: repository.NewTransactionRepository(db),
	}
}

type ITransactionService interface {
}
