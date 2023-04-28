package domain

import (
	"time"

	"github.com/ericlagergren/decimal"
)

type Account struct {
	ID        string
	AccountID string
	Name      string
	Owner     string
	Balance   *decimal.Big
	CreatedAt time.Time
}
