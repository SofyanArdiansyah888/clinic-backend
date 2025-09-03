# Clinic Backend

This is the backend service for the Clinic Management System built with Go Fiber and PostgreSQL.

## 🚨 For Cursor AI Users

**BEFORE making any changes, please read:**
- **[CURSOR_README.md](./CURSOR_README.md)** - Critical instructions for AI
- **[.cursorrules](./.cursorrules)** - Cursor-specific rules
- **[docs/](./docs/)** - Complete project documentation

## 📚 Documentation

- **[📝 Coding Standards](./docs/coding-standards.md)** - Coding patterns and conventions
- **[🏗️ Folder Structure](./docs/folder-structure.md)** - Project organization
- **[🔌 API Documentation](./docs/api-documentation.md)** - API endpoints and patterns
- **[🛠️ Development Guide](./docs/development-guide.md)** - Setup and workflow
- **[🚀 Deployment Guide](./docs/deployment-guide.md)** - Deployment instructions

## 🏗️ Project Structure

```
clinic-backend/
├── .cursor/                 # Cursor IDE configuration
├── .cursorrules            # Cursor AI rules
├── CURSOR_README.md        # AI-specific instructions
├── config/                 # Application configuration
├── docs/                   # Project documentation
├── models/                 # Database models
├── modules/                # Business logic modules
├── routes/                 # Route definitions
├── utils/                  # Utility functions
├── main.go                 # Application entry point
└── go.mod                  # Go module definition
```

## 🚀 Quick Start

1. **Read Documentation First**
   - Check [Development Guide](./docs/development-guide.md) for setup
   - Review [Coding Standards](./docs/coding-standards.md) for patterns

2. **Setup Environment**
   ```bash
   # Install dependencies
   go mod download
   
   # Setup database
   # Follow instructions in development guide
   
   # Run application
   go run main.go
   ```

3. **Follow Standards**
   - Use Bahasa Indonesia for error messages
   - Follow established module patterns
   - Use proper error handling with `utils.Error()` and `utils.Success()`

## 🔧 Technology Stack

- **Framework**: Go Fiber
- **Database**: PostgreSQL with GORM
- **Validation**: go-playground/validator
- **Architecture**: Modular layered architecture

## 📋 Key Features

- RESTful API design
- Modular business logic
- Comprehensive error handling
- Input validation and sanitization
- Database transaction support
- Pagination and filtering
- Security best practices

## 🎯 Architecture Pattern

```
┌─────────────────┐
│   Controllers   │  ← HTTP request/response handling
├─────────────────┤
│    Services     │  ← Business logic implementation
├─────────────────┤
│  Repositories   │  ← Database operations
├─────────────────┤
│     Models      │  ← Data structures and validation
└─────────────────┘
```

## 📞 Support

- Check [documentation](./docs/) first
- Review [coding standards](./docs/coding-standards.md)
- Follow established patterns in existing modules

---

**Remember: Always read documentation before making changes!**
