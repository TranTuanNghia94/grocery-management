# Go API Project

A RESTful API built with Go 1.23.5, Gin, GORM, GORM Gen, golang-migrate, and PostgreSQL, using clean architecture with interfaces.

## Features

- User registration and authentication with JWT
- RESTful API design
- Database migrations
- Environment-based configuration
- Clean architecture with interfaces
- Dependency injection
- Type-safe queries with GORM Gen
- Docker and Docker Compose support

## Clean Architecture

This project follows clean architecture principles with well-defined interfaces between layers:

### Layers
1. **Models**: Data structures and DTOs
2. **Repositories**: Database operations defined by interfaces
3. **Services**: Business logic defined by interfaces
4. **Controllers**: HTTP request handlers defined by interfaces
5. **Routes**: API endpoint definitions

### Benefits of Interface-Based Approach
- **Testability**: Easy to mock dependencies for unit testing
- **Loose Coupling**: Components depend on abstractions, not concrete implementations
- **Flexibility**: Implementations can be swapped without changing the dependent code
- **Clear Boundaries**: Well-defined responsibilities for each layer

## Getting Started

### Local Development

1. Clone the repository
```
git clone https://github.com/yourusername/myapp.git
cd myapp
```

2. Copy the example environment file and update with your values
```
cp .env.example .env
```

3. Install dependencies
```
go mod tidy
```

4. Create the PostgreSQL database
```
createdb myapp
```

5. Run migrations
```
make migrate-up
```

6. Generate type-safe query code
```
make generate
```

7. Run the application
```
make run
```

## Testing

The project uses interfaces to facilitate testing with mocks:

```
go test ./...
```

Example test using mocks:
```go
// Setup mock repository
mockRepo := mocks.NewMockUserRepository()
userService := services.NewUserService(mockRepo, config)

// Test service with mock repository
user, err := userService.CreateUser(input)
```

## API Endpoints

### Public Endpoints

- `POST /api/register` - Register a new user
- `POST /api/login` - Login and get JWT token

### Protected Endpoints (Requires Authentication)

- `GET /api/users` - List all users
- `GET /api/users/:id` - Get a user by ID
- `PUT /api/users/:id` - Update a user
- `DELETE /api/users/:id` - Delete a user