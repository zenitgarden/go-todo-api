package service

import (
	"todo-api/model/web"

	"github.com/gofiber/fiber/v2"
)

type CategoryService interface {
	Create(c *fiber.Ctx, r web.CategoryCreateRequest) web.CategoryResponse
	Update(c *fiber.Ctx, r web.CategoryUpdateRequest) web.CategoryResponse
	Delete(c *fiber.Ctx, r web.CategoryRequest) web.CategoryResponse
	FindById(c *fiber.Ctx, r web.CategoryRequest) web.CategoryResponse
	FindAll(c *fiber.Ctx, owner string) []web.CategoryResponse
}
