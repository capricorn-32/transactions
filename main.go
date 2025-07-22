package main

import (
	"net/http"
	"transactions/config"
	"transactions/db"
	"transactions/handler"
	"transactions/repository"
	"transactions/router"
	"transactions/service"

	_ "github.com/golang-migrate/migrate/v4"
)

func main() {
	logger := config.GetLogger()
	cfg := config.LoadConfig()

	db, err := db.NewDB(cfg)
	if err != nil {
		logger.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	accountRepo := repository.NewAccountRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	accountService := service.NewAccountService(accountRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	h := handler.NewHandler(accountService, transactionService)
	r := router.NewRouter(h)

	logger.Println("Server started at :8080")
	logger.Fatal(http.ListenAndServe(":8080", r))
}
