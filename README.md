# Project Structure

```
transactions/
  ├── go.mod
  ├── go.sum
  ├── main.go
  ├── README.md
  ├── config/
  │     └── config.go
  ├── db/
  │     ├── db.go
  │     └── migrations/
  │           └── 001_init.sql
  ├── models/
  │     └── account.go
  │     └── transaction.go
  ├── repository/
  │     └── account_repository.go
  │     └── transaction_repository.go
  ├── service/
  │     └── account_service.go
  │     └── transaction_service.go
  ├── handler/
  │     └── account_handler.go
  │     └── transaction_handler.go
  ├── router/
  │     └── router.go
  └── tests/
        └── account_handler_test.go
        └── transaction_handler_test.go
```

## Setup, Run, and API Usage

(Instructions to be filled in after code is scaffolded)
