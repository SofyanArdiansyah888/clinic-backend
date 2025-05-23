# CRUD Operation Rules

## Create Operation
- Each create operation must validate input data before processing
- Required fields must be checked and validated
- Return appropriate success/error responses
- Log create operations for audit purposes
- Implement proper error handling

## Read Operation 
- Support pagination for list operations
- Allow filtering and sorting of data
- Implement caching where appropriate
- Return consistent response format
- Handle "not found" scenarios gracefully

## Update Operation
- Validate input data before processing updates
- Check if record exists before updating
- Implement optimistic locking when needed
- Return updated data in response
- Log update operations for audit trail

## Delete Operation
- Verify authorization before deletion
- Implement soft delete when appropriate
- Check dependencies before hard delete
- Return success/failure status
- Log delete operations

## General Rules
- Follow RESTful API conventions
- Use proper HTTP status codes
- Implement request validation
- Handle errors consistently
- Document all endpoints
- Follow security best practices
- Maintain audit logs
- Use proper authentication/authorization
