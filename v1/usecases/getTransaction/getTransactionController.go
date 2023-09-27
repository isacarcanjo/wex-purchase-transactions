package usecases

import (
	"encoding/json"
	"log"
	"net/http"
	"serve/v1/entitys"
	"serve/v1/responses"

	"github.com/gorilla/mux"
)

type GetTransactionController struct {
	IGetTransactionController
}

// Handle @Summary Get transaction
// @Description Get transaction
// @Tags transaction
// @Accept  json
// @Produce  json
// @Param id path string true "Id transaction"
// @Success 200 {object} entitys.TransactionEntity
// @Router /transaction/{id} [get]
func (useCase GetTransactionController) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Add("Content-Type", "application/json")

	if len(id) == 0 {
		responses.Error(w, http.StatusBadRequest, entitys.ErrInvalidId)
		return
	}
	transaction, errTransaction := useCase.Execute(id)
	if errTransaction != nil {
		log.Println(errTransaction.Error())
		responses.Error(w, http.StatusConflict, errTransaction)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if errEncode := json.NewEncoder(w).Encode(transaction); errEncode != nil {
		responses.Error(w, http.StatusInternalServerError, errEncode)
		return
	}
}
