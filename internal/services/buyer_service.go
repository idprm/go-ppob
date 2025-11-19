package services

import "github.com/idprm/go-ppob/internal/domain/repository"

type BuyerService struct {
	buyerRepo repository.IBuyerRepository
}

func NewBuyerService(
	buyerRepo repository.IBuyerRepository,
) *BuyerService {
	return &BuyerService{
		buyerRepo: buyerRepo,
	}
}

type IBuyerService interface {
}
