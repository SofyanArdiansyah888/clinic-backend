# Clinic Backend - Coding Standards

## Overview
Dokumen ini mendefinisikan standar penulisan kode yang konsisten untuk proyek Clinic Backend. Standar ini mengikuti pola yang sudah ada dan memastikan konsistensi di seluruh codebase.

## 📋 Daftar Isi
- [Struktur File dan Penamaan](#struktur-file-dan-penamaan)
- [Konvensi Penamaan](#konvensi-penamaan)
- [Struktur Package](#struktur-package)
- [Format Kode](#format-kode)
- [Error Handling](#error-handling)
- [Response Format](#response-format)
- [Database Operations](#database-operations)
- [Validation](#validation)
- [Documentation](#documentation)
- [Testing](#testing)

## 🏗️ Struktur File dan Penamaan

### Struktur Module
Setiap module harus mengikuti struktur yang konsisten:

```
modules/{module-name}/
├── {module}-controller.go    # HTTP request/response handling
├── {module}-service.go       # Business logic implementation
├── {module}-repository.go    # Database operations
├── {module}-routes.go        # Route definitions
└── [request-models].go       # DTOs (jika diperlukan)
```

### Penamaan File
- **Controllers**: `{module}-controller.go`
- **Services**: `{module}-service.go`
- **Repositories**: `{module}-repository.go`
- **Routes**: `{module}-routes.go`
- **Models**: Gunakan singular form, snake_case untuk multi-word

**Contoh:**
```
antrian-controller.go
antrian-service.go
antrian-repository.go
antrian-routes.go
```

## 🏷️ Konvensi Penamaan

### Package Names
```go
// Gunakan nama package yang sama dengan nama folder
package antrian
package pasien
package staff
```

### Struct Names
```go
// Gunakan PascalCase untuk struct
type Controller struct {
    service IAntrianService
}

type antrianService struct {
    repo IAntrianRepository
}

type antrianRepository struct {
    db *gorm.DB
}
```

### Interface Names
```go
// Gunakan prefix "I" untuk interface
type IAntrianService interface {
    GetAll() ([]models.Antrian, error)
    GetByID(id uint) (*models.Antrian, error)
    Create(data *models.Antrian) error
    Update(id uint, data *models.Antrian) error
    Delete(id uint) error
}

type IAntrianRepository interface {
    FindAll() ([]models.Antrian, error)
    FindByID(id uint) (*models.Antrian, error)
    Create(antrian *models.Antrian) error
    Update(antrian *models.Antrian) error
    Delete(id uint) error
}
```

### Function Names
```go
// Controllers: Gunakan nama yang deskriptif
func (h *Controller) Index(c *fiber.Ctx) error
func (h *Controller) Show(c *fiber.Ctx) error
func (h *Controller) Store(c *fiber.Ctx) error
func (h *Controller) Update(c *fiber.Ctx) error
func (h *Controller) Delete(c *fiber.Ctx) error

// Services: Gunakan nama yang mencerminkan business logic
func (s *antrianService) GetAll() ([]models.Antrian, error)
func (s *antrianService) GetByID(id uint) (*models.Antrian, error)
func (s *antrianService) Create(data *models.Antrian) error

// Repositories: Gunakan nama yang mencerminkan database operations
func (r *antrianRepository) FindAll() ([]models.Antrian, error)
func (r *antrianRepository) FindByID(id uint) (*models.Antrian, error)
func (r *antrianRepository) Create(antrian *models.Antrian) error
```

### Variable Names
```go
// Gunakan camelCase untuk variable
var antrians []models.Antrian
var antrian models.Antrian
var idParam string
var id int
```

### Constants
```go
// Gunakan PascalCase untuk constants
type Status string

const (
    AntrianPasien    Status = "pasien"
    AntrianPerawatan Status = "perawatan"
    AntrianSelesai   Status = "selesai"
    AntrianBatal     Status = "batal"
)
```

## 📦 Struktur Package

### Import Order
```go
import (
    // 1. Standard library imports
    "fmt"
    "strconv"
    "time"

    // 2. Third-party imports
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"

    // 3. Local imports (backend/*)
    "backend/config"
    "backend/models"
    "backend/utils"
)
```

### Package Organization
```go
// Setiap file harus dimulai dengan package declaration
package antrian

// Import statements
import (
    // ... imports
)

// Type definitions
type Controller struct {
    // ...
}

// Constructor functions
func NewAntrianController(service IAntrianService) *Controller {
    return &Controller{service}
}

// Method implementations
func (h *Controller) Index(c *fiber.Ctx) error {
    // ...
}
```

## 📝 Format Kode

### Indentation
- Gunakan tab untuk indentation (bukan spaces)
- Setiap level indentation = 1 tab

### Line Length
- Maksimal 120 karakter per baris
- Jika melebihi, gunakan line break yang logis

### Spacing
```go
// Gunakan spasi setelah keyword
if err != nil {
    return err
}

// Gunakan spasi di sekitar operator
result := a + b

// Jangan gunakan spasi di dalam function call
functionCall(param1, param2)

// Gunakan spasi setelah comma dalam struct literals
struct{
    Field1: value1,
    Field2: value2,
}
```

### Comments
```go
// Gunakan // untuk single line comments
// Gunakan /* */ untuk multi-line comments

// Function comments harus dimulai dengan nama function
// NewAntrianController creates a new antrian controller instance
func NewAntrianController(service IAntrianService) *Controller {
    return &Controller{service}
}

// Method comments harus menjelaskan apa yang dilakukan
// Index returns paginated list of antrian
func (h *Controller) Index(c *fiber.Ctx) error {
    // ...
}
```

## ⚠️ Error Handling

### Error Response Format
```go
// Gunakan utils.Error untuk error responses
if err != nil {
    return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data antrian", err.Error())
}

// Untuk validation errors
if err := c.BodyParser(&antrian); err != nil {
    return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid")
}
```

### Error Checking Pattern
```go
// Selalu check error setelah function call
result, err := someFunction()
if err != nil {
    return err
}

// Untuk multiple error checks, gunakan early return
if err := validateInput(); err != nil {
    return err
}

if err := processData(); err != nil {
    return err
}

return success
```

### Custom Error Messages
```go
// Gunakan pesan error yang deskriptif dalam Bahasa Indonesia
if pasien.NamaPasien == "" {
    return fmt.Errorf("nama pasien wajib diisi")
}

if pasien == nil || antrian == nil {
    return fmt.Errorf("data pasien atau antrian tidak boleh kosong")
}
```

## 📤 Response Format

### Success Response
```go
// Gunakan utils.Success untuk success responses
return utils.Success(c, "Antrian ditemukan", 200, antrian)

// Atau gunakan c.JSON untuk custom response
return c.JSON(fiber.Map{
    "message": "Success",
    "data":    result,
})
```

### Error Response
```go
// Gunakan utils.Error untuk error responses
return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid")

// Atau gunakan c.Status().JSON untuk custom error response
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
    "error": "Data tidak valid",
})
```

## 🗄️ Database Operations

### Model Definition
```go
type Antrian struct {
    ID        uint                 `gorm:"primaryKey" json:"id"`
    IDPasien  uint                 `gorm:"not null" json:"id_pasien"`
    IDStaff   uint                 `gorm:"not null" json:"id_staff"`
    Pasien    Pasien               `gorm:"foreignKey:IDPasien" json:"pasien"`
    Staff     Staff                `gorm:"foreignKey:IDStaff" json:"staff"`
    Tanggal   customtypes.DateTime `gorm:"not null" json:"tanggal"`
    NoAntrian string               `json:"no_antrian"`
    Status    Status               `json:"status" gorm:"default:pasien"`
    CreatedAt time.Time            `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time            `gorm:"autoUpdateTime" json:"updated_at"`
}
```

### Repository Pattern
```go
// Interface definition
type IAntrianRepository interface {
    FindAll() ([]models.Antrian, error)
    FindByID(id uint) (*models.Antrian, error)
    Create(antrian *models.Antrian) error
    Update(antrian *models.Antrian) error
    Delete(id uint) error
}

// Implementation
type antrianRepository struct {
    db *gorm.DB
}

func NewAntrianRepository(db *gorm.DB) IAntrianRepository {
    return &antrianRepository{db: db}
}

func (r *antrianRepository) FindAll() ([]models.Antrian, error) {
    var antrian []models.Antrian
    err := r.db.Find(&antrian).Error
    return antrian, err
}
```

### Transaction Handling
```go
func (s *antrianService) createPasienAntrian(pasien *models.Pasien, antrian *models.Antrian) error {
    // Get DB from repository
    repoImpl, ok := s.repo.(*antrianRepository)
    if !ok || repoImpl.db == nil {
        return fmt.Errorf("repo tidak valid")
    }

    tx := repoImpl.db.Begin()
    if tx.Error != nil {
        return tx.Error
    }

    // Perform operations
    if err := tx.Create(pasien).Error; err != nil {
        tx.Rollback()
        return fmt.Errorf("gagal membuat pasien: %w", err)
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return fmt.Errorf("gagal commit transaksi: %w", err)
    }

    return nil
}
```

## ✅ Validation

### Input Validation
```go
// Gunakan go-playground/validator untuk struct validation
type CreateAntrianRequest struct {
    IDPasien  uint   `json:"id_pasien" validate:"required"`
    IDStaff   uint   `json:"id_staff" validate:"required"`
    NoAntrian string `json:"no_antrian" validate:"required"`
}

// Validation in controller
func (h *Controller) Store(c *fiber.Ctx) error {
    var antrian models.Antrian
    if err := c.BodyParser(&antrian); err != nil {
        return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid")
    }
    
    // Additional validation
    if antrian.IDPasien == 0 {
        return utils.Error(c, fiber.StatusBadRequest, "ID Pasien wajib diisi")
    }
    
    // ... rest of the logic
}
```

### Custom Validation
```go
// Gunakan FormatValidationError untuk validation errors
func FormatValidationError(err error) map[string]string {
    errors := make(map[string]string)
    for _, e := range err.(validator.ValidationErrors) {
        errors[e.Field()] = fmt.Sprintf("Field '%s' is %s", e.Field(), e.ActualTag())
    }
    return errors
}
```

## 📚 Documentation

### Function Documentation
```go
// NewAntrianController creates a new antrian controller instance
// Parameters:
//   - service: The antrian service interface
// Returns:
//   - *Controller: A new controller instance
func NewAntrianController(service IAntrianService) *Controller {
    return &Controller{service}
}

// Index returns paginated list of antrian
// Parameters:
//   - c: Fiber context
// Returns:
//   - error: Any error that occurred during processing
func (h *Controller) Index(c *fiber.Ctx) error {
    // Implementation
}
```

### Package Documentation
```go
// Package antrian provides functionality for managing patient queues
// 
// This package includes:
// - Controllers for HTTP request handling
// - Services for business logic
// - Repositories for data access
// - Routes for API endpoints
package antrian
```

## 🧪 Testing

### Test File Structure
```go
// File: antrian_service_test.go
package antrian

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// Test function naming: Test{Struct}_{Method}
func TestAntrianService_GetByID(t *testing.T) {
    // Setup
    mockRepo := &MockAntrianRepository{}
    service := NewAntrianService(mockRepo)
    
    // Test data
    expectedAntrian := &models.Antrian{ID: 1, NoAntrian: "ANT001"}
    
    // Mock expectations
    mockRepo.On("FindByID", uint(1)).Return(expectedAntrian, nil)
    
    // Execute
    result, err := service.GetByID(1)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expectedAntrian, result)
    mockRepo.AssertExpectations(t)
}
```

### Test Naming Conventions
```go
// Test{Struct}_{Method}_{Scenario}
func TestAntrianService_GetByID_Success(t *testing.T) { }
func TestAntrianService_GetByID_NotFound(t *testing.T) { }
func TestAntrianService_Create_InvalidData(t *testing.T) { }
```

## 🔧 Best Practices

### Dependency Injection
```go
// Gunakan constructor pattern untuk dependency injection
type Controller struct {
    service IAntrianService
}

func NewAntrianController(service IAntrianService) *Controller {
    return &Controller{service}
}
```

### Interface Segregation
```go
// Definisikan interface yang spesifik
type IAntrianService interface {
    GetAll() ([]models.Antrian, error)
    GetByID(id uint) (*models.Antrian, error)
    Create(data *models.Antrian) error
    Update(id uint, data *models.Antrian) error
    Delete(id uint) error
}
```

### Error Wrapping
```go
// Gunakan fmt.Errorf dengan %w untuk error wrapping
if err := tx.Create(pasien).Error; err != nil {
    tx.Rollback()
    return fmt.Errorf("gagal membuat pasien: %w", err)
}
```

### Logging
```go
// Gunakan structured logging untuk debugging
import "log"

func (s *antrianService) Create(data *models.Antrian) error {
    log.Printf("Creating antrian with data: %+v", data)
    
    err := s.repo.Create(data)
    if err != nil {
        log.Printf("Failed to create antrian: %v", err)
        return err
    }
    
    log.Printf("Antrian created successfully with ID: %d", data.ID)
    return nil
}
```

## 📋 Checklist

### Sebelum Commit
- [ ] Kode mengikuti format yang ditentukan
- [ ] Error handling sudah diimplementasi dengan benar
- [ ] Response format konsisten
- [ ] Validation sudah ditambahkan
- [ ] Comments sudah ditambahkan untuk fungsi kompleks
- [ ] Tests sudah ditulis (jika applicable)
- [ ] Naming conventions sudah diikuti
- [ ] Import order sudah benar

### Code Review
- [ ] Struktur kode mengikuti pattern yang ada
- [ ] Error messages dalam Bahasa Indonesia
- [ ] Response format konsisten
- [ ] Database operations menggunakan transaction jika diperlukan
- [ ] Validation sudah mencukupi
- [ ] Performance considerations sudah diperhatikan

---

**Catatan**: Standar ini harus diikuti oleh semua developer yang berkontribusi pada proyek. Setiap perubahan pada standar ini harus didiskusikan dengan tim dan didokumentasikan dengan baik.
