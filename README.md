# Transactions Service

A simple Go REST API for managing accounts and transactions, using PostgreSQL as the backend database.

## Project Structure

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
  │           ├── 20250710183538_transactions.sql
  │           └── 20250710183542_accounts.sql
  ├── models/
  │     ├── account.go
  │     ├── transaction.go
  │     └── money.go
  ├── repository/
  │     ├── account_repository.go
  │     └── transaction_repository.go
  ├── service/
  │     ├── account_service.go
  │     └── transaction_service.go
  ├── handler/
  │     ├── account_handler.go
  │     └── transaction_handler.go
  ├── router/
  │     └── router.go
  └── tests/
        ├── account_handler_test.go
        └── transaction_handler_test.go
```

## Installation

1. **Clone the repository:**
   ```sh
   git clone git@github.com:capricorn-32/transactions.git
   cd transactions
   ```
2. **Install Go (>=1.24.4) and PostgreSQL.**
3. **Install Task (taskfile.dev) for task runner:**
   ```sh
   # On Linux (with Homebrew)
   brew install go-task/tap/go-task
   # Or see https://taskfile.dev/installation/
   ```
4. **Install Go dependencies:**
   ```sh
   go mod download
   ```

## Setup

1. **Configure environment variables (optional, defaults shown):**
   - `DB_USER` (default: `postgres`)
   - `DB_PASSWORD` (default: `postgres`)
   - `DB_NAME` (default: `postgres`)
   - `DB_HOST` (default: `localhost`)
   - `DB_PORT` (default: `5433`)

   You can set these in your shell or in a `.env` file (not committed).

2. **Create and migrate the database:**
   ```sh
   task migrate
   # Or, to reset (drop, create, migrate):
   task reset
   ```

## Running the Application

- **Run the server:**
  ```sh
  task run
  # Or manually:
  go run main.go
  ```
- The server will start on `http://localhost:8080` by default.

## Running Tests

```sh
task test
# Or manually:
go test ./... -v
```

## API Endpoints

### Create Account
- **POST** `/accounts`
- **Request Body:**
  ```json
  {
    "account_id": 1,
    "initial_balance": "100.00"
  }
  ```
- **Response:**
  - `201 Created` on success:
    ```json
    { "success": true, "message": "account created successfully" }
    ```
  - `400 Bad Request` on error:
    ```json
    { "success": false, "error": "error message" }
    ```

### Get Account
- **GET** `/accounts/{account_id}`
- **Response:**
  - `200 OK` on success:
    ```json
    { "success": true, "data": { "account_id": 1, "balance": "100.00" } }
    ```
  - `400 Bad Request` or `404 Not Found` on error:
    ```json
    { "success": false, "error": "error message" }
    ```

### Submit Transaction
- **POST** `/transactions`
- **Request Body:**
  ```json
  {
    "source_account_id": 1,
    "destination_account_id": 2,
    "amount": "50.00"
  }
  ```
- **Response:**
  - `201 Created` on success:
    ```json
    { "success": true, "message": "transaction submitted successfully" }
    ```
  - `400 Bad Request` on error:
    ```json
    { "success": false, "error": "error message" }
    ```

## Assumptions & Notes

- The application expects a running PostgreSQL instance accessible with the provided credentials.
- Default database connection is to `localhost:5433` with user/password `postgres`.
- All balances and amounts are strings representing decimal numbers (e.g., "100.00").
- No authentication is implemented; all endpoints are public for demo/testing purposes.
- The database schema is managed via SQL files in `db/migrations/`.
- The application uses the [Task](https://taskfile.dev) runner for common dev tasks.


