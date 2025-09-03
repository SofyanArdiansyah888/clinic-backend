# 🚨 CURSOR AI - READ THIS FIRST

## 📚 MANDATORY: Read Documentation Before Making Changes

**BEFORE writing ANY code or making ANY changes, you MUST read these documents:**

### 🔥 CRITICAL DOCUMENTATION (READ FIRST)
1. **📝 [Coding Standards](./docs/coding-standards.md)** - **ALL coding patterns and conventions**
2. **🏗️ [Folder Structure](./docs/folder-structure.md)** - Project organization and architecture
3. **🔌 [API Documentation](./docs/api-documentation.md)** - API patterns and response formats
4. **🛠️ [Development Guide](./docs/development-guide.md)** - Setup and workflow

## 🎯 Key Points for AI

### Language Requirements
- **ALWAYS use Bahasa Indonesia** for error messages and user-facing text
- **NEVER use English** for error messages or user feedback

### Framework Requirements
- **ONLY use Go Fiber** (NO Gin or other frameworks)
- **ONLY use libraries already in go.mod**
- **Use GORM for database operations**
- **Use go-playground/validator for validation**

### Code Patterns (MUST FOLLOW)
```go
// Error handling - ALWAYS use utils.Error
return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid")

// Success responses - ALWAYS use utils.Success
return utils.Success(c, "Berhasil dibuat", 201, data)

// Interface naming - ALWAYS use "I" prefix
type IAntrianService interface { ... }

// File naming - ALWAYS follow pattern
antrian-controller.go
antrian-service.go
antrian-repository.go
antrian-routes.go
```

### Module Structure (MANDATORY)
```
modules/{module-name}/
├── {module}-controller.go    # HTTP handling
├── {module}-service.go       # Business logic
├── {module}-repository.go    # Database operations
└── {module}-routes.go        # Route definitions
```

## 📋 Before Writing Code Checklist

- [ ] ✅ Read `/docs/coding-standards.md`
- [ ] ✅ Check `/docs/folder-structure.md`
- [ ] ✅ Review existing similar modules
- [ ] ✅ Follow naming conventions
- [ ] ✅ Use proper error handling
- [ ] ✅ Implement validation
- [ ] ✅ Use Bahasa Indonesia for messages

## 🚫 What NOT to Do

- ❌ Don't use different frameworks (only Go Fiber)
- ❌ Don't add new external libraries
- ❌ Don't use English for error messages
- ❌ Don't ignore existing patterns
- ❌ Don't skip validation
- ❌ Don't forget error handling

## ✅ What TO Do

- ✅ Always read documentation first
- ✅ Follow established patterns
- ✅ Use Bahasa Indonesia for messages
- ✅ Implement proper validation
- ✅ Handle errors correctly
- ✅ Maintain consistency
- ✅ Update documentation when needed

## 🔍 Code Review Focus

When suggesting code changes, ensure:
1. Follows coding standards from `/docs/coding-standards.md`
2. Maintains consistency with existing codebase
3. Uses proper error handling patterns
4. Implements validation correctly
5. Uses Bahasa Indonesia for messages
6. Follows repository-service-controller pattern

---

## 🚨 REMINDER FOR AI

**DOCUMENTATION IS THE SOURCE OF TRUTH**

- **ALWAYS** check `/docs/coding-standards.md` before writing code
- **ALWAYS** follow existing patterns in the codebase
- **ALWAYS** use Bahasa Indonesia for user messages
- **ALWAYS** maintain consistency with existing modules

**If you're unsure about anything, READ THE DOCUMENTATION FIRST!**
