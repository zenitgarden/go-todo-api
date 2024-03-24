package app

import (
	"todo-api/controller"
	"todo-api/repository"
	"todo-api/routes"
	"todo-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func InitApi(api fiber.Router) {
	db := NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, categoryRepository, db, validate)
	todoController := controller.NewTodoController(todoService)

	routes.UserRouter(api, userController)
	routes.CategoryRouter(api, categoryController)
	routes.TodoRouter(api, todoController)
}
