package account

import (
	"account/core/domain"
	"github.com/jmoiron/sqlx"
)

type RepositoryMysql struct {
	dbClient *sqlx.DB
}

func NewAccountRepositoryMysql(dbClient *sqlx.DB) *RepositoryMysql {
	return &RepositoryMysql{dbClient}
}

func (a *RepositoryMysql) FindByID(ID string) (*domain.Account, error) {

	var accountDB AccountDBMysql
	row := a.dbClient.QueryRow(findByID, ID)

	err := row.Scan(
		&accountDB.ID,
		&accountDB.AccountID,
		&accountDB.Name,
		&accountDB.Owner,
		&accountDB.Balance,
		&accountDB.CreatedAt)
	if err != nil {
		return nil, err
	}

	return accountDB.toDomain(), nil
}
