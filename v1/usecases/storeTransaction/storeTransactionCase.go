package usecases

import (
	"serve/v1/entitys"
	"serve/v1/repositories"
)

type CreateUserCase struct {
	repo repositories.ITransactionRepository
}

func NewService(r repositories.ITransactionRepository) *CreateUserCase {
	return &CreateUserCase{
		repo: r,
	}
}

func (s *CreateUserCase) Execute(input StoreTransactionInput) (entitys.TransactionEntity, error) {
	transaction, err := entitys.NewTransaction(input.Description, input.TransactionDate, input.PurchaseAmount)
	if err != nil {
		return transaction, err
	}
	newTransaction, createErr := s.repo.Create(transaction)
	if createErr != nil {
		return entitys.TransactionEntity{}, createErr
	}
	return newTransaction, createErr
}
