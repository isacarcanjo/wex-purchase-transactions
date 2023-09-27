# wex-purchase-transactions
![Generic badge](https://img.shields.io/badge/mantainer-isacarcanjo-black.svg)
![Generic badge](https://img.shields.io/badge/Golang-1.18.0-blue.svg)
![Generic badge](https://img.shields.io/badge/version-1.0.0-gree.svg)

## Description
Service responsible to purchase transactions and retrieve a Purchase Transaction in a Specified Country’s
Currency.

## Tech Stack
1. [Golang 1.18+](https://go.dev/blog/go1.18)

## Methods
1. Clean architecture
2. SOLID

## Routes
* /v1
  * [POST /v1/transaction](http://localhost:5000/v1/swagger/index.html#/transaction/post_transaction)
  
  * [GET /v1/transaction/id](http://localhost:5000/v1/swagger/index.html#/transaction/get_transaction__id_)
  

## Local Variables

If not exist, you should make .env file
```
PORT=5000
VERSION=v1
TRANSACTION_COLLECTION=transaction
```

## Installing
The easier way is run this

```
go get ./...
``` 

or you can see the [documentation](https://go.dev/doc/modules/managing-dependencies)

## Run Application

How to run the application

```
go run main.go
```

## Documentation
http://localhost:5000/v1/swagger/index.html

## Requirements

You need to develop a RESTful API that stores a Purchase Transactions and retrieve a Purchase Transaction in a Specified Country’s
Currency.

- ✅ Store a new transactions - description (max 50), transaction date, and a purchase amount in United States dollars (only positive value).
- ✅ Retrieve a transaction - transaction id required
- ✅ Documentation + Setup - We need to run your code based on README and
use the API based on the documentation.
