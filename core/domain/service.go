package domain

import "github.com/ericlagergren/decimal"

func (a *Account) WithDraw(money *decimal.Big) {
	a.Balance = new(decimal.Big).Sub(a.Balance, money)
}

func (a *Account) Deposit(money *decimal.Big) {
	a.Balance = new(decimal.Big).Add(a.Balance, money)
}
