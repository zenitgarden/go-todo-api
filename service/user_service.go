package service

import (
	"todo-api/model/web"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Register(c *fiber.Ctx, r web.UserCreateRequest) web.UserResponse
	Update(c *fiber.Ctx, r web.UserUpdateRequest) web.UserResponse
	Login(c *fiber.Ctx, r web.LoginRequest) web.LoginResponse
}
