package implementations

import (
	"encoding/json"
	"fmt"
	"serve/v1/entitys"

	scribble "github.com/nanobox-io/golang-scribble"
)

type SimDBProvider struct {
	TransactionCollection string
}

func (s SimDBProvider) connect() (*scribble.Driver, error) {
	db, errDb := scribble.New("./", nil)
	if errDb != nil {
		fmt.Println("Error", errDb)
		return nil, errDb
	}
	return db, nil
}

func (u SimDBProvider) Create(e entitys.TransactionEntity) (entitys.TransactionEntity, error) {
	db, errDb := u.connect()
	if errDb != nil {
		fmt.Println("Error", errDb)
	}

	if err := db.Write(u.TransactionCollection, e.Id, e); err != nil {
		fmt.Println("Error", err)
		return entitys.TransactionEntity{}, nil
	}
	return e, nil
}

func (u SimDBProvider) Find(id string) (entitys.TransactionEntity, error) {
	db, errDb := u.connect()
	if errDb != nil {
		fmt.Println("Error", errDb)
	}
	transaction := entitys.TransactionEntity{}
	if err := db.Read(u.TransactionCollection, id, &transaction); err != nil {
		fmt.Println("Error", err)
		return transaction, entitys.ErrNotFound
	}
	return transaction, nil
}

func (u SimDBProvider) Update(e entitys.TransactionEntity) (entitys.TransactionEntity, error) {
	db, errDb := u.connect()
	if errDb != nil {
		return entitys.TransactionEntity{}, nil
	}

	if errDb != nil {
		fmt.Println("Error", errDb)
	}
	if err := db.Write(u.TransactionCollection, e.Id, e); err != nil {
		fmt.Println("Error", err)
		return entitys.TransactionEntity{}, nil
	}
	return e, nil
}

func (u SimDBProvider) FindAll() ([]entitys.TransactionEntity, error) {
	db, errDb := u.connect()
	var transactions []entitys.TransactionEntity
	if errDb != nil {
		fmt.Println("Error", errDb)
		return transactions, nil
	}
	records, err := db.ReadAll(u.TransactionCollection)
	if err != nil {
		fmt.Println("Error", err)
		return transactions, nil
	}
	for _, f := range records {
		transactionFound := entitys.TransactionEntity{}
		if errUnmarshal := json.Unmarshal([]byte(f), &transactionFound); errUnmarshal != nil {
			fmt.Println("Error", errUnmarshal)
			return transactions, errUnmarshal
		}
		transactions = append(transactions, transactionFound)
	}
	return transactions, nil
}
