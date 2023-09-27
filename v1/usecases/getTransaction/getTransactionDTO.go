package usecases

import (
	"serve/v1/entitys"
)

type IGetTransactionController interface {
	IGetTransactionCase
}

type IGetTransactionCase interface {
	Execute(id string) (entitys.TransactionEntity, error)
}

type GetTransactionInput struct {
	Id string `json:"id"`
}
