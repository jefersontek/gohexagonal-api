package api

import (
	v1 "account/api/v1"
	"account/core/ports"
	"account/core/usecase"
	"account/env"
	"account/repository/account"
	"account/util"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func Start() {

	var accountPort ports.IAccountPort

	if env.Environment.Database == "mongo" {
		clientDB, err := util.OpenMongoConnection(env.Environment.MongoURI)
		if err != nil {
			os.Exit(1)
		}
		accountPort = account.NewAccountRepositoryMongo(clientDB)
		defer clientDB.Disconnect(context.TODO())
	} else {
		clientDB, err := util.OpenMySQLConnection(env.Environment.MysqlDSN)
		if err != nil {
			os.Exit(1)
		}
		accountPort = account.NewAccountRepositoryMysql(clientDB)
		defer clientDB.Close()
	}

	accountUseCase := usecase.NewAccountUseCase(accountPort)

	router := mux.NewRouter()
	baseRouter := router.PathPrefix("/").Subrouter()

	v1.Map(baseRouter, accountUseCase)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", env.Environment.Port),
		Handler: router,
	}

	printRoutes(router)

	// inicia o servidor em uma goroutine separada para não bloquear
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			os.Exit(1)
		}
	}()

	// channel que bloqueia a função start até que seja recebido um sigkill ou sigterm
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	// à partir daqui, o sinal de termination foi recebido, então fechamos os recursos e matamos o servidor
	shutdownContext, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	httpServer.Shutdown(shutdownContext)
}

func printRoutes(router *mux.Router) {
	// print routes
	fmt.Println("ROUTES:")
	builder := strings.Builder{}

	err := router.Walk(func(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {

		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			builder.WriteString(" - ")
			builder.WriteString(pathTemplate)
		}
		methods, err := route.GetMethods()
		if err == nil {
			builder.WriteString(" -- ")
			builder.WriteString(strings.Join(methods, ", "))
		}
		builder.WriteString("\n")

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(builder.String())
	fmt.Println()
}
