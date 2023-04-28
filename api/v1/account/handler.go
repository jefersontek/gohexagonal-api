package account

import (
	"account/core/usecase"
	"account/util"
	"net/http"
)

func get(core *usecase.AccountUseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		AccountID := r.URL.Query().Get("id")
		account, _ := core.GetByAccountID(AccountID)

		util.ReplyJSON(w, http.StatusOK, toResponse(account))
		return
	}
}
