package usecases

import (
	config2 "serve/v1/config"
	"serve/v1/providers/implementations"
	repositories "serve/v1/repositories/implements"
)

func MakeGetTransaction() GetTransactionController {
	config := config2.GetDotEnv()
	db := implementations.SimDBProvider{TransactionCollection: config.TransactionCollection}
	userRepo := repositories.NewTransactionRepository(db)
	makeGetTransactionCase := NewService(userRepo)
	makeController := GetTransactionController{IGetTransactionController: makeGetTransactionCase}
	return makeController
}
