package usecases

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"serve/v1/responses"
)

type StoreTransactionController struct {
	IStoreTransactionController
}

// Handle @Summary Create transaction
// @Description Create transaction
// @Tags transaction
// @Accept  json
// @Produce  json
// @Param transaction body StoreTransactionInput true "Transaction"
// @Success 200 {object} entitys.TransactionEntity
// @Router /transaction [post]
func (useCase StoreTransactionController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	var input StoreTransactionInput
	if errUnmarshalL := json.Unmarshal(body, &input); errUnmarshalL != nil {
		responses.Error(w, http.StatusUnprocessableEntity, errUnmarshalL)
		return
	}
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	newTransaction, errTransaction := useCase.Execute(input)
	if errTransaction != nil {
		log.Println(errTransaction.Error())
		responses.Error(w, http.StatusConflict, errTransaction)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if errEncode := json.NewEncoder(w).Encode(newTransaction); errEncode != nil {
		responses.Error(w, http.StatusInternalServerError, errEncode)
		return
	}
}
