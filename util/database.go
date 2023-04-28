package util

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenMySQLConnection(dsn string) (*sqlx.DB, error) {

	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	return db, nil
}

func OpenMongoConnection(mongoURI string) (*mongo.Client, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	return client, nil
}
