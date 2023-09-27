package entitys

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("transaction not found")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot Be Deleted")

// ErrInvalidDescription ErrInvalidName name is invalid
var ErrInvalidDescription = errors.New("description is Invalid")

// ErrInvalidTransactionDate ErrInvalidName name is invalid
var ErrInvalidTransactionDate = errors.New("transaction date is Invalid")

// ErrInvalidPurchaseAmount ErrInvalidName name is invalid
var ErrInvalidPurchaseAmount = errors.New("purchase amount date is Invalid")

//ErrInvalidId id is invalid
var ErrInvalidId = errors.New("id is Invalid")
