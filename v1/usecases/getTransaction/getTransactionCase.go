package usecases

import (
	"serve/v1/entitys"
	"serve/v1/repositories"
)

type GetUserCase struct {
	repo repositories.ITransactionRepository
}

func NewService(r repositories.ITransactionRepository) *GetUserCase {
	return &GetUserCase{
		repo: r,
	}
}

func (s *GetUserCase) Execute(id string) (entitys.TransactionEntity, error) {
	transaction, getErr := s.repo.Find(id)
	if getErr != nil {
		return entitys.TransactionEntity{}, getErr
	}
	return transaction, nil
}
