package entitys

import (
	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type TransactionEntity struct {
	Id              string  `json:"id"`
	Description     string  `json:"description"`
	TransactionDate string  `json:"transaction_date"`
	PurchaseAmount  float64 `json:"purchase_amount"`
}

// NewTransaction Create a transaction
func NewTransaction(description, transactionDate string, purchaseAmount float64) (TransactionEntity, error) {
	//ac := accounting.Accounting{Symbol: "$", Precision: 2}
	s := TransactionEntity{}
	if len(description) == 0 || len(description) > 50 {
		return s, ErrInvalidDescription
	}
	if len(transactionDate) == 0 {
		return s, ErrInvalidTransactionDate
	}
	if purchaseAmount < 0 {
		return s, ErrInvalidPurchaseAmount
	}
	s.Id = uuid.New().String()
	s.Description = strings.TrimSpace(description)
	s.TransactionDate = transactionDate
	s.PurchaseAmount = purchaseAmount
	//fmt.Println(ac.FormatMoney(purchaseAmount))
	//fmt.Println(helpers.RoundFloat(purchaseAmount, 2))
	//fmt.Println(purchaseAmount)

	return s, nil
}

func (s TransactionEntity) Update(newS TransactionEntity) TransactionEntity {
	newTransaction := s
	if len(newS.Description) > 0 {
		newTransaction.Description = cases.Title(language.Und, cases.NoLower).String(strings.ToLower(newS.Description))
	}
	if len(newS.TransactionDate) > 0 {
		newTransaction.TransactionDate = strings.TrimSpace(newS.TransactionDate)
	}
	if newS.PurchaseAmount >= 0 {
		newTransaction.PurchaseAmount = newS.PurchaseAmount
	}
	return newTransaction
}
