package service

import (
	"todo-api/model/web"

	"github.com/gofiber/fiber/v2"
)

type TodoService interface {
	Create(c *fiber.Ctx, r web.TodoCreateRequest) web.TodoResponse
	Update(c *fiber.Ctx, r web.TodoUpdateRequest) web.TodoResponse
	Delete(c *fiber.Ctx, r web.TodoRequest) web.TodoDeleteResponse
	FindById(c *fiber.Ctx, r web.TodoRequest) web.TodoResponse
	FindAll(c *fiber.Ctx, owner string, qs web.TQueryString) []web.TodoResponse
}
