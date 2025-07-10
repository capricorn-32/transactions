# Transactions Service

A simple Go REST API for managing accounts and transactions, built with PostgreSQL.

## 🚀 Quick Start

### Prerequisites
- **Go** (>=1.24.4)
- **PostgreSQL** (running instance)
- **Task** (taskfile.dev) - [Install here](https://taskfile.dev/installation/)

### One-Command Setup

```bash
# Clone and setup everything
git clone git@github.com:capricorn-32/transactions.git
cd transactions
task setup
task migrate
task run
```

That's it! Your API is now running at `http://localhost:8080`.

## 📋 Available Commands

Run `task` to see all available commands:

| Command | Description |
|---------|-------------|
| `task` | Show all available tasks |
| `task setup` | Install all dependencies |
| `task run` | Start the application |
| `task build` | Build the application |
| `task test` | Run all tests |
| `task migrate` | Run database migrations |
| `task reset` | Reset database (rollback + migrate) |

## 🔧 Configuration

### Environment Variables (Optional)

Set these if you need custom database settings:

```bash
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_HOST=localhost
export DB_PORT=5433
export DB_NAME=postgres
```

**Defaults work out of the box** - no configuration needed if using standard PostgreSQL setup.

## 🗄️ Database Setup

The project uses PostgreSQL with automatic migrations:

```bash
# First time setup
task migrate

# Reset database (useful for development)
task reset
```

## 🧪 Testing

```bash
# Run all tests
task test
```

## 📦 Building

```bash
# Build for production
task build
```

The binary will be created at `bin/transactions`.

## 🔌 API Endpoints

### Create Account
```bash
POST /accounts
Content-Type: application/json

{
  "account_id": 1,
  "initial_balance": "100.00"
}
```

### Get Account
```bash
GET /accounts/{account_id}
```

### Submit Transaction
```bash
POST /transactions
Content-Type: application/json

{
  "source_account_id": 1,
  "destination_account_id": 2,
  "amount": "50.00"
}
```

## 🛠️ Development Workflow

### Typical Development Session

```bash
# 1. Reset database to clean state
task reset

# 2. Start the application
task run

# 3. In another terminal, run tests
task test

# 4. Build for deployment
task build
```

### Database Management

```bash
# Apply new migrations
task migrate

# Reset database (useful during development)
task reset
```

## 📁 Project Structure

```
transactions/
├── main.go                 # Application entry point
├── Taskfile.yml           # Task runner configuration
├── go.mod                 # Go module dependencies
├── config/
│   └── config.go         # Configuration management
├── db/
│   ├── db.go             # Database connection
│   └── migrations/       # Database migration files
├── models/
│   ├── account.go        # Account model
│   ├── transaction.go    # Transaction model
│   └── money.go          # Money handling utilities
├── repository/
│   ├── account_repository.go    # Account data access
│   └── transaction_repository.go # Transaction data access
├── service/
│   ├── account_service.go       # Account business logic
│   └── transaction_service.go   # Transaction business logic
├── handler/
│   ├── account_handler.go       # Account HTTP handlers
│   └── transaction_handler.go   # Transaction HTTP handlers
├── router/
│   └── router.go               # HTTP routing
└── tests/
    ├── account_handler_test.go
    └── transaction_handler_test.go
```

## 🐛 Troubleshooting

### Common Issues

**Database connection errors:**
- Ensure PostgreSQL is running
- Check that the database exists and is accessible
- Verify environment variables if using custom settings

**Migration errors:**
- Run `task setup` to ensure golang-migrate is installed
- Check that PostgreSQL is running and accessible

**Build errors:**
- Ensure Go version >=1.24.4
- Run `task setup` to install dependencies

## 📝 Notes

- All balances and amounts are strings representing decimal numbers (e.g., "100.00")
- No authentication implemented - all endpoints are public for demo/testing
- Database migrations are managed with golang-migrate
- The application uses Task for common development operations

