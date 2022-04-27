package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivek-101/banking/dto"
	"github.com/vivek-101/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_Id := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customer_Id
		account, appErr := h.service.NewAccount(request)
		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.Message)
		} else {
			WriteResponse(w, http.StatusCreated, account)
		}
	}

}
