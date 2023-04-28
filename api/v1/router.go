package v1

import (
	"account/api/v1/account"
	"account/core/usecase"
	"github.com/gorilla/mux"
)

// Map mapeia as rotas da API
func Map(r *mux.Router, core *usecase.AccountUseCase) {
	account.Map(r.PathPrefix("/v1").Subrouter(), core)
}
