# Transactions Service

A simple Go REST API for managing accounts and transactions, using PostgreSQL as the backend database.

## Project Structure

```
transactions/
  ├── go.mod
  ├── go.sum
  ├── main.go
  ├── README.md
  ├── Taskfile.yml
  ├── config/
  │     └── config.go
  ├── db/
  │     ├── db.go
  │     └── migrations/
  │           ├── 000001_accounts.up.sql
  │           ├── 000001_accounts.down.sql
  │           ├── 000002_transactions.up.sql
  │           └── 000002_transactions.down.sql
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

## Prerequisites

- **Go** (>=1.24.4)
- **PostgreSQL** (running instance)
- **Task** (taskfile.dev) - task runner
- **golang-migrate** - for database migrations

## Installation

1. **Clone the repository:**
   ```sh
   git clone git@github.com:capricorn-32/transactions.git
   cd transactions
   ```

2. **Install Go dependencies:**
   ```sh
   go mod download
   ```

3. **Install Task (taskfile.dev) for task runner:**
   ```sh
   # On Linux (with Homebrew)
   brew install go-task/tap/go-task
   
   # On macOS (with Homebrew)
   brew install go-task/tap/go-task
   
   # On Windows (with Chocolatey)
   choco install go-task
   
   # Or see https://taskfile.dev/installation/ for other methods
   ```

4. **Install golang-migrate:**
   ```sh
   # On Linux/macOS
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   
   # Or download from https://github.com/golang-migrate/migrate/releases
   ```

## Configuration

### Environment Variables

The following environment variables can be configured (defaults shown):

- `DB_USER` (default: `postgres`)
- `DB_PASSWORD` (default: `postgres`)
- `DB_NAME` (default: `postgres`)
- `DB_HOST` (default: `localhost`)
- `DB_PORT` (default: `5433`)

You can set these in your shell or in a `.env` file (not committed to version control).

### Database Setup

1. **Ensure PostgreSQL is running** and accessible with the configured credentials
2. **Run database migrations:**
   ```sh
   task migrate
   ```

## Available Tasks

This project uses [Task](https://taskfile.dev) for common development operations. Run `task` to see all available commands.

### Development Tasks

- **`task`** - Show all available tasks with descriptions and usage
- **`task setup`** - Install all dependencies (Go modules and golang-migrate)
- **`task run`** - Run the Go application
- **`task build`** - Build the Go application (outputs to `bin/transactions`)
- **`task test`** - Run all Go tests with verbose output

### Database Tasks

- **`task migrate`** - Run all up migrations using golang-migrate
- **`task reset`** - Rollback all migrations and then run all up migrations (useful for development)

## Quick Start
1. **Install dependencies:**
   ```sh
   task setup
   ```

2. **Migrate the database:**
   ```sh
   task migrate
   ```

3. **Start the application:**
   ```sh
   task run
   ```
   The server will start on `http://localhost:8080` by default.

4. **Run tests:**
   ```sh
   task test
   ```

5. **Build the application:**
   ```sh
   task build
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

## Development Workflow

### Typical Development Session

1. **Start with a clean database state:**
   ```sh
   task reset
   ```

2. **Run the application:**
   ```sh
   task run
   ```

3. **In another terminal, run tests:**
   ```sh
   task test
   ```

4. **Build for deployment:**
   ```sh
   task build
   ```

### Database Management

- **Apply new migrations:**
  ```sh
  task migrate
  ```

- **Reset database (useful during development):**
  ```sh
  task reset
  ```

## Assumptions & Notes

- The application expects a running PostgreSQL instance accessible with the provided credentials
- Default database connection is to `localhost:5433` with user/password `postgres`
- All balances and amounts are strings representing decimal numbers (e.g., "100.00")
- No authentication is implemented; all endpoints are public for demo/testing purposes
- The database schema is managed via SQL files in `db/migrations/`
- The application uses the [Task](https://taskfile.dev) runner for common dev tasks
- Database migrations use [golang-migrate](https://github.com/golang-migrate/migrate) for version control

## Troubleshooting

### Common Issues

1. **Database connection errors:**
   - Ensure PostgreSQL is running
   - Verify environment variables are set correctly
   - Check that the database exists and is accessible

2. **Migration errors:**
   - Ensure golang-migrate is installed with postgres support
   - Check that the database URL is correct
   - Verify migration files are in the correct format

3. **Build errors:**
   - Ensure Go version >=1.24.4
   - Run `go mod download` to install dependencies
   - Check that all required environment variables are set

