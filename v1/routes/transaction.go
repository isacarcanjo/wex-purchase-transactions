package routes

import (
	getTransaction "serve/v1/usecases/getTransaction"
	storeTransaction "serve/v1/usecases/storeTransaction"
)

func GetTransactionRouters() []Router {
	var TransactionRouters = []Router{
		{
			URI:                   "/transaction",
			Method:                "POST",
			DomainFunc:            storeTransaction.MakeStoreTransaction().Handle,
			RequireAuthentication: true,
		},
		{
			URI:                   "/transaction/{id}",
			Method:                "GET",
			DomainFunc:            getTransaction.MakeGetTransaction().Handle,
			RequireAuthentication: false,
		},
	}
	return TransactionRouters
}
