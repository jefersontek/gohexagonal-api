package usecase

import (
	"account/core/domain"
	"account/core/ports"
)

type IAccountUseCase interface {
	GetByAccountID(ID string) (*domain.Account, error)
}

type AccountUseCase struct {
	accountRepository ports.IAccountPort
}

func NewAccountUseCase(accountRepository ports.IAccountPort) *AccountUseCase {
	return &AccountUseCase{accountRepository: accountRepository}
}

func (a *AccountUseCase) GetByAccountID(ID string) (*domain.Account, error) {
	account, err := a.accountRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return account, nil
}
