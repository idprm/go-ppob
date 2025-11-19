package services

import "github.com/idprm/go-ppob/internal/domain/repository"

type APIKeyService struct {
	apiKeyRepo repository.IAPIKeyRepository
}

func NewAPIKeyService(apiKeyRepo repository.IAPIKeyRepository) *APIKeyService {
	return &APIKeyService{
		apiKeyRepo: apiKeyRepo,
	}
}

type IAPIKeyService interface{}
