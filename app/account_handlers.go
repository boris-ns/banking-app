package app

import (
	"banking-app/dto"
	"banking-app/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah *AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
	var req dto.NewAccountRequestDto
	req.CustomerId = mux.Vars(r)["id"]

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		account, appErr := ah.service.NewAccount(req)

		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			writeResponse(w, http.StatusCreated, account.ToNewAccountResponseDto())
		}
	}
}
