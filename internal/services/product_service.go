package services

import "github.com/idprm/go-ppob/internal/domain/repository"

type ProductService struct {
	productRepo repository.IProductRepository
}

func NewProductService(
	productRepo repository.IProductRepository,
) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

type IProductService interface {
}
