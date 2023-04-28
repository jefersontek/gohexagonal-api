package account

import (
	"account/core/domain"
	"account/util"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountDBMysql struct {
	ID        sql.NullString  `db:"id"`
	AccountID sql.NullString  `db:"account_id"`
	Name      sql.NullString  `db:"name"`
	Owner     sql.NullString  `db:"owner"`
	Balance   sql.NullFloat64 `db:"balance"`
	CreatedAt mysql.NullTime  `db:"created_at"`
}

type AccountDBMongo struct {
	ID        string             `bson:"id"`
	AccountID string             `bson:"account_id"`
	Name      string             `bson:"name"`
	Owner     string             `bson:"owner"`
	Balance   float64            `bson:"balance"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}

func (a *AccountDBMysql) toDomain() *domain.Account {
	return &domain.Account{
		ID:        a.ID.String,
		AccountID: a.AccountID.String,
		Name:      a.Name.String,
		Owner:     a.Owner.String,
		Balance:   util.FloatToDecimal(a.Balance.Float64),
		CreatedAt: a.CreatedAt.Time,
	}
}

func (a *AccountDBMongo) toDomain() *domain.Account {
	return &domain.Account{
		ID:        a.ID,
		AccountID: a.AccountID,
		Name:      a.Name,
		Owner:     a.Owner,
		Balance:   util.FloatToDecimal(a.Balance),
		CreatedAt: a.CreatedAt.Time(),
	}
}
