package services

import "github.com/idprm/go-ppob/internal/domain/repository"

type SellerService struct {
	sellerRepo repository.ISellerRepository
}

func NewSellerService(
	sellerRepo repository.ISellerRepository,
) *SellerService {
	return &SellerService{
		sellerRepo: sellerRepo,
	}
}

type ISellerService interface {
}
