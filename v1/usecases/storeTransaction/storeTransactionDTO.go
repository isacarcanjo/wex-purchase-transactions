package usecases

import (
	"serve/v1/entitys"
)

type IStoreTransactionController interface {
	IStoreTransactionCase
}

type IStoreTransactionCase interface {
	Execute(StoreTransactionInput) (entitys.TransactionEntity, error)
}

type StoreTransactionInput struct {
	Description     string  `json:"description"`
	TransactionDate string  `json:"transaction_date"`
	PurchaseAmount  float64 `json:"purchase_amount"`
}
