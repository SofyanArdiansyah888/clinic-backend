package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = fmt.Sprintf("Field '%s' is %s", e.Field(), e.ActualTag())
	}
	return errors
}

// Success returns a JSON success response
func Success(c *fiber.Ctx, message string, responseCode int, data ...interface{}) error {
	resp := fiber.Map{
		"message": message,
	}
	if len(data) > 0 {
		resp["data"] = data[0]
	}
	return c.Status(responseCode).JSON(resp)
}

// Error returns a JSON error response
func Error(c *fiber.Ctx, code int, message string, detail ...string) error {
	resp := fiber.Map{
		"error": message,
	}
	if len(detail) > 0 {
		resp["message"] = detail[0]
	}
	return c.Status(code).JSON(resp)
}
