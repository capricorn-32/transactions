package main

import (
	"net/http"
	"transactions/config"
	"transactions/db"
	"transactions/handler"
	"transactions/repository"
	"transactions/router"
	"transactions/service"
)

func main() {
	logger := config.GetLogger()
	cfg := config.LoadConfig()
	postgresDB, err := db.NewDB(cfg)
	if err != nil {
		logger.Fatalf("failed to connect to db: %v", err)
	}
	defer postgresDB.Close()

	accountRepo := repository.NewAccountRepository(postgresDB)
	transactionRepo := repository.NewTransactionRepository(postgresDB)

	accountService := service.NewAccountService(accountRepo)
	transactionService := service.NewTransactionService(postgresDB, transactionRepo)

	h := handler.NewHandler(accountService, transactionService)
	r := router.NewRouter(h)

	logger.Println("Server started at :8080")
	logger.Fatal(http.ListenAndServe(":8080", r))
}
