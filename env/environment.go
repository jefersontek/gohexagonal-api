package env

import (
	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
	"log"
)

type environment struct {
	Environment string `env:"ENVIRONMENT" default:"development"`
	Port        int    `env:"PORT" default:"9000"`
	Database    string `env:"DATABASE" default:"mysql"`
	MysqlDSN    string `env:"MYSQL_DSN" required:"true"`
	MongoURI    string `env:"MONGO_URI" required:"true"`
}

// Config ...
var Environment environment

// InitConfigs inicializa as configurações
func InitConfigs() {

	// carrega o .env caso exista
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("No .env file found")
	}

	// bind env vars
	if err := env.Set(&Environment); err != nil {
		log.Fatal(err)
	}
}
