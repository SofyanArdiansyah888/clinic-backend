# Clinic Backend - API Documentation

## Overview
This document provides comprehensive API documentation for the Clinic Backend system. All endpoints follow RESTful conventions and return consistent JSON responses.

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication
Most endpoints require authentication. Include the following header:
```
Authorization: Bearer <your-jwt-token>
```

## Response Format

### Success Response
```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": {},
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "total_pages": 10
  },
  "errors": null
}
```

### Error Response
```json
{
  "success": false,
  "message": "Operation failed",
  "data": null,
  "errors": [
    {
      "field": "email",
      "message": "Invalid email format"
    }
  ]
}
```

## HTTP Status Codes
- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `500` - Internal Server Error

## API Endpoints

### Authentication

#### POST /auth/login
User login endpoint.

**Request Body:**
```json
{
  "username": "admin",
  "password": "password123"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@clinic.com",
      "role": "admin"
    }
  }
}
```

### User Management

#### GET /users
Get all users with pagination.

**Query Parameters:**
- `page` (int): Page number (default: 1)
- `limit` (int): Items per page (default: 10)
- `search` (string): Search term
- `role` (string): Filter by role

**Response:**
```json
{
  "success": true,
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": 1,
      "username": "admin",
      "email": "admin@clinic.com",
      "role": "admin",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 50,
    "total_pages": 5
  }
}
```

#### POST /users
Create a new user.

**Request Body:**
```json
{
  "username": "newuser",
  "email": "user@clinic.com",
  "password": "password123",
  "role": "staff"
}
```

#### PUT /users/{id}
Update user information.

**Request Body:**
```json
{
  "username": "updateduser",
  "email": "updated@clinic.com",
  "role": "admin"
}
```

#### DELETE /users/{id}
Delete a user.

### Patient Management

#### GET /patients
Get all patients with pagination and filtering.

**Query Parameters:**
- `page` (int): Page number
- `limit` (int): Items per page
- `search` (string): Search by name or ID
- `status` (string): Filter by status

**Response:**
```json
{
  "success": true,
  "message": "Patients retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "phone": "+628123456789",
      "email": "john@example.com",
      "birth_date": "1990-01-01",
      "address": "Jl. Example No. 123",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

#### POST /patients
Create a new patient.

**Request Body:**
```json
{
  "name": "Jane Doe",
  "phone": "+628123456789",
  "email": "jane@example.com",
  "birth_date": "1995-05-15",
  "address": "Jl. Example No. 456",
  "gender": "female"
}
```

### Appointment Management

#### GET /appointments
Get all appointments.

**Query Parameters:**
- `page` (int): Page number
- `limit` (int): Items per page
- `patient_id` (int): Filter by patient
- `date` (string): Filter by date (YYYY-MM-DD)
- `status` (string): Filter by status

#### POST /appointments
Create a new appointment.

**Request Body:**
```json
{
  "patient_id": 1,
  "treatment_id": 1,
  "appointment_date": "2024-01-15T10:00:00Z",
  "notes": "Regular checkup"
}
```

### Treatment Management

#### GET /treatments
Get all treatments.

**Response:**
```json
{
  "success": true,
  "message": "Treatments retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Regular Checkup",
      "description": "Standard health checkup",
      "duration": 30,
      "price": 150000,
      "category": "general"
    }
  ]
}
```

#### POST /treatments
Create a new treatment.

**Request Body:**
```json
{
  "name": "Dental Cleaning",
  "description": "Professional dental cleaning service",
  "duration": 60,
  "price": 300000,
  "category": "dental"
}
```

### Product/Item Management

#### GET /products
Get all products with pagination.

**Query Parameters:**
- `page` (int): Page number
- `limit` (int): Items per page
- `category` (string): Filter by category
- `search` (string): Search by name or code

#### POST /products
Create a new product.

**Request Body:**
```json
{
  "name": "Vitamin C",
  "code": "VIT-C-001",
  "category": "supplement",
  "price": 50000,
  "stock": 100,
  "description": "Vitamin C supplement"
}
```

### Queue Management

#### GET /queue
Get current queue status.

**Response:**
```json
{
  "success": true,
  "message": "Queue status retrieved",
  "data": {
    "current_number": "A001",
    "total_waiting": 5,
    "estimated_wait_time": 25
  }
}
```

#### POST /queue
Add patient to queue.

**Request Body:**
```json
{
  "patient_id": 1,
  "treatment_id": 1,
  "priority": "normal"
}
```

### Stock Management

#### GET /stock
Get stock information.

**Query Parameters:**
- `product_id` (int): Filter by product
- `category` (string): Filter by category

#### POST /stock/movement
Record stock movement.

**Request Body:**
```json
{
  "product_id": 1,
  "type": "in",
  "quantity": 50,
  "notes": "Stock replenishment"
}
```

## Error Handling

### Validation Errors
When input validation fails, the API returns a 400 status with detailed error messages:

```json
{
  "success": false,
  "message": "Validation failed",
  "errors": [
    {
      "field": "email",
      "message": "Invalid email format"
    },
    {
      "field": "phone",
      "message": "Phone number is required"
    }
  ]
}
```

### Authentication Errors
When authentication fails:

```json
{
  "success": false,
  "message": "Unauthorized access",
  "data": null,
  "errors": null
}
```

### Not Found Errors
When a resource is not found:

```json
{
  "success": false,
  "message": "Resource not found",
  "data": null,
  "errors": null
}
```

## Rate Limiting
API endpoints are rate-limited to prevent abuse:
- 100 requests per minute for authenticated users
- 10 requests per minute for unauthenticated users

## Pagination
List endpoints support pagination with the following parameters:
- `page`: Page number (starts from 1)
- `limit`: Number of items per page (max 100)

## Filtering and Search
Many endpoints support filtering and search:
- `search`: Text search across relevant fields
- `status`: Filter by status
- `date`: Filter by date range
- `category`: Filter by category

## File Upload
For file uploads (e.g., patient documents), use multipart/form-data:

```
POST /patients/{id}/documents
Content-Type: multipart/form-data

file: [binary file data]
description: "Medical certificate"
```

## WebSocket Endpoints
Real-time updates are available via WebSocket:

```
ws://localhost:8080/ws/queue
```

Subscribe to queue updates and receive real-time notifications.

## Testing
Use the provided Postman collection or curl commands to test the API endpoints.

### Example curl commands:

```bash
# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password123"}'

# Get patients
curl -X GET http://localhost:8080/api/v1/patients \
  -H "Authorization: Bearer <token>"

# Create patient
curl -X POST http://localhost:8080/api/v1/patients \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","phone":"+628123456789"}'
```

## Versioning
API versioning is handled through URL paths:
- Current version: `/api/v1/`
- Future versions: `/api/v2/`, `/api/v3/`, etc.

## Deprecation Policy
- Deprecated endpoints will be marked with a deprecation header
- Deprecated endpoints will be supported for at least 6 months
- Migration guides will be provided for breaking changes

This documentation should be updated whenever new endpoints are added or existing ones are modified.
