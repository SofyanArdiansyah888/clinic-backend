# Clinic Backend - Development Guide

## Overview
This guide provides comprehensive instructions for setting up and developing the Clinic Backend project. The project uses Go with Fiber framework and follows a modular architecture pattern.

## Prerequisites

### Required Software
- **Go**: Version 1.21 or higher
- **PostgreSQL**: Version 13 or higher
- **Git**: For version control
- **Air**: For live reloading (optional but recommended)

### Optional Tools
- **Postman**: For API testing
- **Docker**: For containerized development
- **VS Code/Cursor**: Recommended IDE with Go extensions

## Project Setup

### 1. Clone the Repository
```bash
git clone <repository-url>
cd clinic-backend
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Database Setup
Create a PostgreSQL database and update the configuration:

```bash
# Create database
createdb clinic_db

# Or using psql
psql -U postgres
CREATE DATABASE clinic_db;
```

### 4. Environment Configuration
Create a `.env` file in the root directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=clinic_db
DB_SSL_MODE=disable

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=localhost

# JWT Configuration
JWT_SECRET=your_jwt_secret_key
JWT_EXPIRY=24h

# Environment
ENV=development
```

### 5. Run Database Migrations
The project uses GORM auto-migration. Run the application to automatically create tables:

```bash
go run main.go
```

## Development Workflow

### Starting the Development Server

#### Using Air (Recommended)
```bash
# Install Air if not already installed
go install github.com/cosmtrek/air@latest

# Start development server with live reload
air
```

#### Using Go Run
```bash
go run main.go
```

#### Using Go Build
```bash
go build -o clinic-backend
./clinic-backend
```

### Project Structure Understanding

#### Adding New Features
1. **Create/Update Models** (`/models`)
   - Define database schema
   - Add validation tags
   - Define relationships

2. **Create Module** (`/modules`)
   - Follow the established pattern
   - Implement controller, service, repository, and routes

3. **Update Routes** (`/routes`)
   - Add new module routes to main routes file

4. **Add Tests**
   - Unit tests for business logic
   - Integration tests for API endpoints

### Code Organization

#### Module Structure
Each module should follow this structure:

```
modules/your-module/
├── your-module-controller.go    # HTTP handling
├── your-module-service.go       # Business logic
├── your-module-repository.go    # Database operations
├── your-module-routes.go        # Route definitions
└── [request-models].go          # DTOs (if needed)
```

#### File Naming Conventions
- **Controllers**: `{module}-controller.go`
- **Services**: `{module}-service.go`
- **Repositories**: `{module}-repository.go`
- **Routes**: `{module}-routes.go`
- **Models**: Use singular form, snake_case for multi-word

### Database Development

#### Creating New Models
1. Create model file in `/models`
2. Define struct with GORM tags
3. Add validation tags
4. Define relationships

Example:
```go
type Patient struct {
    BaseModel
    Name      string    `json:"name" gorm:"not null" validate:"required"`
    Phone     string    `json:"phone" gorm:"uniqueIndex" validate:"required"`
    Email     string    `json:"email" validate:"email"`
    BirthDate time.Time `json:"birth_date"`
    Address   string    `json:"address"`
    Gender    string    `json:"gender" validate:"oneof=male female"`
}
```

#### Database Migrations
- Use GORM auto-migration
- Run `go run main.go` to apply migrations
- Check database schema after changes

### API Development

#### Creating New Endpoints
1. **Controller**: Handle HTTP requests/responses
2. **Service**: Implement business logic
3. **Repository**: Handle database operations
4. **Routes**: Define API endpoints

#### Response Format
Always use the standard response format:

```go
type Response struct {
    Success    bool        `json:"success"`
    Message    string      `json:"message"`
    Data       interface{} `json:"data"`
    Pagination *Pagination `json:"pagination,omitempty"`
    Errors     []Error     `json:"errors,omitempty"`
}
```

#### Error Handling
Implement proper error handling:

```go
func (c *Controller) Create(ctx *fiber.Ctx) error {
    // Parse request
    var request CreateRequest
    if err := ctx.BodyParser(&request); err != nil {
        return ctx.Status(400).JSON(Response{
            Success: false,
            Message: "Invalid request body",
            Errors:  []Error{{Field: "body", Message: err.Error()}},
        })
    }
    
    // Validate request
    if err := validator.New().Struct(request); err != nil {
        return ctx.Status(400).JSON(Response{
            Success: false,
            Message: "Validation failed",
            Errors:  parseValidationErrors(err),
        })
    }
    
    // Process request
    result, err := c.service.Create(request)
    if err != nil {
        return ctx.Status(500).JSON(Response{
            Success: false,
            Message: "Failed to create resource",
        })
    }
    
    return ctx.Status(201).JSON(Response{
        Success: true,
        Message: "Resource created successfully",
        Data:    result,
    })
}
```

### Testing

#### Unit Tests
Create unit tests for business logic:

```go
func TestPatientService_Create(t *testing.T) {
    // Setup
    mockRepo := &MockPatientRepository{}
    service := NewPatientService(mockRepo)
    
    // Test data
    request := CreatePatientRequest{
        Name:  "John Doe",
        Phone: "+628123456789",
    }
    
    // Mock expectations
    mockRepo.On("Create", mock.Anything).Return(&Patient{ID: 1}, nil)
    
    // Execute
    result, err := service.Create(request)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, uint(1), result.ID)
    mockRepo.AssertExpectations(t)
}
```

#### Integration Tests
Test API endpoints:

```go
func TestPatientController_Create(t *testing.T) {
    // Setup
    app := fiber.New()
    controller := NewPatientController(service)
    app.Post("/patients", controller.Create)
    
    // Test data
    requestBody := `{"name":"John Doe","phone":"+628123456789"}`
    
    // Execute
    req := httptest.NewRequest("POST", "/patients", strings.NewReader(requestBody))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)
    
    // Assert
    assert.Equal(t, 201, resp.StatusCode)
    
    var response Response
    json.NewDecoder(resp.Body).Decode(&response)
    assert.True(t, response.Success)
}
```

### Code Quality

#### Linting and Formatting
```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run

# Run vet
go vet ./...
```

#### Pre-commit Hooks
Set up pre-commit hooks to ensure code quality:

```bash
# Install pre-commit
pip install pre-commit

# Install hooks
pre-commit install
```

### Debugging

#### Logging
Use structured logging for debugging:

```go
import "log"

func (s *Service) Create(request CreateRequest) (*Model, error) {
    log.Printf("Creating resource with data: %+v", request)
    
    // ... business logic
    
    log.Printf("Resource created successfully with ID: %d", result.ID)
    return result, nil
}
```

#### Database Debugging
Enable GORM debug mode in development:

```go
// In config/database.go
if os.Getenv("ENV") == "development" {
    db = db.Debug()
}
```

### Performance Optimization

#### Database Queries
- Use proper indexing
- Implement pagination
- Optimize N+1 queries
- Use database transactions for complex operations

#### Caching
Implement caching for frequently accessed data:

```go
// Example with Redis
func (r *Repository) GetByID(id uint) (*Model, error) {
    // Check cache first
    if cached, err := r.cache.Get(fmt.Sprintf("model:%d", id)); err == nil {
        return cached, nil
    }
    
    // Query database
    var model Model
    if err := r.db.First(&model, id).Error; err != nil {
        return nil, err
    }
    
    // Cache result
    r.cache.Set(fmt.Sprintf("model:%d", id), &model, time.Hour)
    
    return &model, nil
}
```

### Security Best Practices

#### Input Validation
Always validate input data:

```go
type CreateRequest struct {
    Name  string `json:"name" validate:"required,min=2,max=100"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"required,min=0,max=150"`
}
```

#### Authentication
Implement proper authentication middleware:

```go
func AuthMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Get("Authorization")
        if token == "" {
            return c.Status(401).JSON(Response{
                Success: false,
                Message: "Authorization header required",
            })
        }
        
        // Validate token
        claims, err := validateToken(token)
        if err != nil {
            return c.Status(401).JSON(Response{
                Success: false,
                Message: "Invalid token",
            })
        }
        
        // Set user context
        c.Locals("user", claims)
        return c.Next()
    }
}
```

### Deployment Preparation

#### Environment Configuration
Create environment-specific configurations:

```bash
# Development
cp .env.example .env.development

# Production
cp .env.example .env.production
```

#### Build Optimization
```bash
# Build for production
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o clinic-backend .

# Build with specific version
go build -ldflags="-X main.Version=1.0.0" -o clinic-backend .
```

## Troubleshooting

### Common Issues

#### Database Connection
- Check database credentials
- Ensure PostgreSQL is running
- Verify network connectivity

#### Port Conflicts
- Change server port in configuration
- Check if port is already in use

#### Import Errors
- Run `go mod tidy`
- Check module dependencies
- Verify import paths

### Getting Help
- Check existing documentation
- Review error logs
- Consult team members
- Create issue tickets for bugs

## Best Practices Summary

1. **Follow established patterns** for consistency
2. **Write tests** for all business logic
3. **Validate input** at all entry points
4. **Handle errors** gracefully
5. **Document code** for complex logic
6. **Use meaningful names** for variables and functions
7. **Keep functions small** and focused
8. **Implement proper logging** for debugging
9. **Follow security best practices**
10. **Optimize performance** where needed

This development guide should be updated as the project evolves and new patterns emerge.
