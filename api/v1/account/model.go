package account

import (
	"account/core/domain"
	"account/util"
	"time"
)

type Response struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func toResponse(account *domain.Account) Response {
	return Response{
		ID:        account.ID,
		AccountID: account.AccountID,
		Name:      account.Name,
		Owner:     account.Owner,
		Balance:   util.DecimalToFloat(account.Balance),
		CreatedAt: account.CreatedAt,
	}
}
