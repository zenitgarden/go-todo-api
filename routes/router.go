package routes

import (
	"todo-api/controller"
	"todo-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router, userController controller.UserController) {
	router.Post("/users/register", userController.Register)
	router.Post("/users/login", userController.Login)

	router.Use(middleware.AuthMiddleware())
	router.Put("/users", userController.Update)
}

func CategoryRouter(router fiber.Router, categoryController controller.CategoryController) {
	router.Post("/categories", categoryController.Save)
	router.Put("/categories/:categoryId", categoryController.Update)
	router.Delete("/categories/:categoryId", categoryController.Delete)
	router.Get("/categories/:categoryId", categoryController.FindById)
	router.Get("/categories", categoryController.FindAll)
}

func TodoRouter(router fiber.Router, todoController controller.TodoController) {
	router.Post("/todos", todoController.Save)
	router.Put("/todos/:todoId", todoController.Update)
	router.Delete("/todos/:todoId", todoController.Delete)
	router.Get("/todos/:todoId", todoController.FindById)
	router.Get("/todos", todoController.FindAll)
}
