package user

import (
    "backend/config"
    "github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
    repository := NewUserRepository(config.DB)
    service := NewUserService(repository)
    controller := NewUserController(service)

    group := api.Group("/user")
    group.Get("/", controller.Index)
    group.Get("/:id", controller.Show)
    group.Post("/", controller.Store)
    group.Put("/:id", controller.Update)
    group.Delete("/:id", controller.Delete)
}