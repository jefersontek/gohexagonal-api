package ports

import "account/core/domain"

type IAccountPort interface {
	FindByID(ID string) (*domain.Account, error)
}
