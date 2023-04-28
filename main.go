package main

import (
	"account/api"
	"account/env"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// inicializa configuração
	env.InitConfigs()
	api.Start()
}
