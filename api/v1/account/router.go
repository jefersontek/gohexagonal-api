package account

import (
	"account/core/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func Map(r *mux.Router, core *usecase.AccountUseCase) {
	r.HandleFunc("/accounts", get(core)).Methods(http.MethodGet)
	r.HandleFunc("/accounts/", get(core)).Methods(http.MethodGet)
}
