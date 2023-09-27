package repositories

import (
	"serve/v1/entitys"
)

// ITransactionRepository Repository interface
type ITransactionRepository interface {
	Reader
	Writer
}

type TransactionDBConn interface {
	Connect() (interface{}, error)
}

//Reader interface
type Reader interface {
	Find(id string) (entitys.TransactionEntity, error)
	FindAll() ([]entitys.TransactionEntity, error)
}

//Writer Transaction writer
type Writer interface {
	Create(e entitys.TransactionEntity) (entitys.TransactionEntity, error)
	Update(e entitys.TransactionEntity) (entitys.TransactionEntity, error)
}
