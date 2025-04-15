# Grocery Management System

This project is a simple Grocery Management System built with Go and PostgreSQL. It allows users to manage grocery items and user accounts through a RESTful API.

## Project Structure

```
grocery-management
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── api
│   │   ├── handlers.go      # HTTP handler functions
│   │   └── routes.go        # API routes setup
│   ├── models
│   │   ├── grocery.go       # Grocery item model
│   │   └── user.go          # User model
│   ├── db
│   │   └── postgres.go      # Database connection and interactions
│   └── config
│       └── config.go        # Application configuration management
├── migrations
│   ├── 001_create_users_table.sql  # Migration for users table
│   └── 002_create_groceries_table.sql  # Migration for groceries table
├── scripts
│   └── seed.go              # Database seeding script
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
├── docker-compose.yml       # Docker configuration
├── .env.example             # Example environment variables
└── README.md                # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd grocery-management
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Set up the database:**
   - Create a PostgreSQL database and update the connection details in the `.env` file.

4. **Run migrations:**
   - Use the migration scripts in the `migrations` folder to set up the database schema.

5. **Start the application:**
   ```
   go run cmd/server/main.go
   ```

## Usage

- The API provides endpoints for managing groceries and users. Refer to the `internal/api/routes.go` file for a list of available endpoints and their usage.

## License

This project is licensed under the MIT License.