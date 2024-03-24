package main

import (
	"todo-api/app"
	"todo-api/app/exception"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	server := fiber.New(fiber.Config{ErrorHandler: exception.ErrorHandler})
	server.Use(recover.New())
	api := server.Group("/api")
	app.InitApi(api)
	server.Use(exception.NotFoundHandler)

	server.Listen("localhost:8080")
}
