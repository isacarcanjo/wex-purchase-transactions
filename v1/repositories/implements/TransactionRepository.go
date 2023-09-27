package repositories

import (
	"serve/v1/entitys"
	"serve/v1/repositories"
)

type TransactionRepository struct {
	db repositories.ITransactionRepository
}

//NewTransactionRepository create new repository
func NewTransactionRepository(db repositories.ITransactionRepository) *TransactionRepository {
	return &TransactionRepository{db}
}

//Create transaction
func (r *TransactionRepository) Create(e entitys.TransactionEntity) (entitys.TransactionEntity, error) {
	transaction, err := r.db.Create(e)
	if err != nil {
		return entitys.TransactionEntity{}, err
	}
	return transaction, nil
}

//Find transaction
func (r *TransactionRepository) Find(id string) (entitys.TransactionEntity, error) {
	transaction, err := r.db.Find(id)
	if err != nil {
		return entitys.TransactionEntity{}, err
	}
	return transaction, nil
}

//Update transaction
func (r *TransactionRepository) Update(e entitys.TransactionEntity) (entitys.TransactionEntity, error) {
	transaction, err := r.db.Update(e)
	if err != nil {
		return entitys.TransactionEntity{}, err
	}
	return transaction, nil
}

//FindAll transactions
func (r *TransactionRepository) FindAll() ([]entitys.TransactionEntity, error) {
	var transactions []entitys.TransactionEntity
	transactions, err := r.db.FindAll()
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
