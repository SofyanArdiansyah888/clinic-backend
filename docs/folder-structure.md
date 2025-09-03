# Clinic Backend - Folder Structure Documentation

## Overview
This document describes the folder structure and organization of the Clinic Backend project, which follows a modular architecture pattern using Go Fiber framework.

## Root Directory Structure

```
clinic-backend/
├── .cursor/                 # Cursor IDE configuration
├── .trae/                   # Trae configuration files
├── config/                  # Application configuration
├── docs/                    # Project documentation
├── models/                  # Database models and entities
├── modules/                 # Business logic modules
├── routes/                  # Route definitions
├── tmp/                     # Temporary files
├── utils/                   # Utility functions and helpers
├── .air.toml               # Air live reload configuration
├── .gitignore              # Git ignore rules
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── main.go                  # Application entry point
└── postgres.yml            # PostgreSQL configuration
```

## Detailed Structure

### `/config` - Configuration Files
Contains application configuration and setup files.

```
config/
└── database.go             # Database connection and configuration
```

**Purpose**: Centralized configuration management for database connections, environment variables, and application settings.

### `/docs` - Documentation
Contains project documentation and guides.

```
docs/
├── folder-structure.md      # This file - folder structure documentation
├── api-documentation.md     # API endpoints documentation
├── development-guide.md     # Development setup and guidelines
└── deployment-guide.md      # Deployment instructions
```

**Purpose**: Comprehensive documentation for developers, maintainers, and stakeholders.

### `/models` - Database Models
Contains GORM models representing database tables and relationships.

```
models/
├── customTypes/            # Custom data types and validators
│   ├── date-only.go        # Date-only custom type
│   └── date-time.go        # DateTime custom type
├── antrian.go              # Queue/Appointment queue model
├── appointment.go          # Appointment model
├── appointment-treatment.go # Appointment-treatment relationship
├── bank.go                 # Bank information model
├── barang.go               # Product/Item model
├── cabang.go               # Branch model
├── konversi-stok.go        # Stock conversion model
├── konversi-stok-detail.go # Stock conversion detail model
├── login-history.go        # User login history model
├── membership.go           # Membership model
├── models.go               # Base model definitions
├── monthly-sequence.go     # Monthly sequence generator model
├── pasien.go               # Patient model
├── pembelian.go            # Purchase model
├── pembelian-detail.go     # Purchase detail model
├── penjualan.go            # Sales model
├── penjualan-detail.go     # Sales detail model
├── perawatan.go            # Treatment/Care model
├── perusahaan.go           # Company model
├── produksi-barang.go      # Product production model
├── produksi-barang-detail.go # Product production detail model
├── promo.go                # Promotion model
├── staff.go                # Staff/Employee model
├── stok-movement.go        # Stock movement model
├── stok-opname.go          # Stock opname model
├── supplier.go             # Supplier model
├── template_concern.go     # Template concern model
├── treatment.go            # Treatment model
├── user.go                 # User model
└── voucher.go              # Voucher model
```

**Purpose**: Define database schema, relationships, and data validation rules using GORM.

### `/modules` - Business Logic Modules
Contains modular business logic organized by domain/feature.

```
modules/
├── antrian/                # Queue management module
├── appointment/            # Appointment management module
├── bank/                   # Bank management module
├── barang/                 # Product/Item management module
├── cabang/                 # Branch management module
├── generateNumber/         # Number generation utilities
├── kartuStok/             # Stock card management module
├── konversiBarang/        # Product conversion module
├── membership/            # Membership management module
├── pasien/                # Patient management module
├── pembelianBarang/       # Product purchase module
├── penjualanBarang/       # Product sales module
├── perawatan/             # Treatment/Care management module
├── perusahaan/            # Company management module
├── produksiBarang/        # Product production module
├── promo/                 # Promotion management module
├── staff/                 # Staff management module
├── stokOpname/            # Stock opname module
├── supplier/              # Supplier management module
├── templateConcern/       # Template concern module
├── transaksiBarang/       # Product transaction module
├── treatment/             # Treatment management module
├── user/                  # User management module
└── voucher/               # Voucher management module
```

#### Module Structure Pattern
Each module follows a consistent structure:

```
module-name/
├── module-controller.go    # HTTP request/response handling
├── module-service.go       # Business logic implementation
├── module-repository.go    # Database operations
├── module-routes.go        # Route definitions
└── [request-models].go     # Request/response DTOs (if needed)
```

**Example - Antrian Module**:
```
antrian/
├── antrian-controller.go   # Handle HTTP requests for queue operations
├── antrian-service.go      # Business logic for queue management
├── antrian-repository.go   # Database operations for queue data
└── antrian-routes.go       # Define queue-related API routes
```

**Purpose**: Organize business logic by domain, making the codebase maintainable and scalable.

### `/routes` - Route Definitions
Contains centralized route configuration.

```
routes/
└── routes.go              # Main route configuration and middleware setup
```

**Purpose**: Centralized route management, middleware configuration, and API endpoint organization.

### `/utils` - Utility Functions
Contains reusable utility functions and helpers.

```
utils/
├── generator.go           # ID and number generation utilities
├── helper.go             # General helper functions
└── pagination.go         # Pagination utilities
```

**Purpose**: Provide reusable utility functions across the application.

### `/tmp` - Temporary Files
Contains temporary files generated during development or runtime.

**Purpose**: Store temporary files, logs, and cache data.

## Architecture Principles

### 1. Modular Design
- Each business domain has its own module
- Modules are self-contained with their own controllers, services, and repositories
- Clear separation of concerns

### 2. Layered Architecture
- **Controller Layer**: Handle HTTP requests/responses
- **Service Layer**: Implement business logic
- **Repository Layer**: Handle database operations
- **Model Layer**: Define data structures

### 3. Dependency Management
- Use Go modules for dependency management
- Only use approved libraries (see go.mod)
- Follow Go best practices for imports and package organization

### 4. Configuration Management
- Centralized configuration in `/config`
- Environment-based configuration
- Database configuration separation

## File Naming Conventions

### Models
- Use singular form: `user.go`, `patient.go`
- Use snake_case for multi-word names: `stock_movement.go`

### Modules
- Use camelCase: `userManagement/`, `stockOpname/`
- Module files follow pattern: `{module}-{layer}.go`

### Controllers
- Pattern: `{module}-controller.go`
- Handle HTTP requests and responses
- Input validation and error handling

### Services
- Pattern: `{module}-service.go`
- Business logic implementation
- Orchestrate operations between repositories

### Repositories
- Pattern: `{module}-repository.go`
- Database operations
- Query optimization and data access

### Routes
- Pattern: `{module}-routes.go`
- API endpoint definitions
- Middleware configuration

## Best Practices

### 1. Module Organization
- Keep related functionality together
- Maintain consistent file structure across modules
- Use clear, descriptive module names

### 2. Code Organization
- Follow Go conventions for package naming
- Use meaningful file and function names
- Implement proper error handling

### 3. Documentation
- Document complex business logic
- Maintain up-to-date API documentation
- Include setup and deployment guides

### 4. Testing
- Write unit tests for business logic
- Implement integration tests for API endpoints
- Test error scenarios and edge cases

## Development Workflow

### Adding New Features
1. Create or update models in `/models`
2. Implement business logic in appropriate module
3. Add routes in module's route file
4. Update main routes configuration
5. Add tests and documentation

### Module Creation
1. Create module directory in `/modules`
2. Implement controller, service, repository, and routes
3. Follow established naming conventions
4. Add module routes to main routes file
5. Update documentation

### Database Changes
1. Update or create models in `/models`
2. Use GORM auto-migration
3. Test database operations
4. Update documentation if needed

## Security Considerations

- Input validation in controllers
- Authentication and authorization middleware
- SQL injection prevention through GORM
- Proper error handling without information leakage
- Rate limiting and request validation

This folder structure promotes maintainability, scalability, and clear separation of concerns while following Go and Fiber framework best practices.
