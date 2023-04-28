package account

import (
	"account/core/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryMongo struct {
	dbClient *mongo.Client
}

func NewAccountRepositoryMongo(dbClient *mongo.Client) *RepositoryMongo {
	return &RepositoryMongo{dbClient}
}

func (a *RepositoryMongo) FindByID(ID string) (*domain.Account, error) {

	var accountDB AccountDBMongo
	filterAccountID := bson.D{{"account_id", ID}}

	err := a.dbClient.Database("account").
		Collection("accounts").
		FindOne(context.TODO(), filterAccountID).
		Decode(&accountDB)
	if err != nil {
		return nil, err
	}

	return accountDB.toDomain(), nil
}
