package controller

import (
	"encoding/json"
	"todo-api/helper"
	"todo-api/model/web"
	"todo-api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(c *fiber.Ctx) error {
	body := c.Body()
	req := new(web.UserCreateRequest)
	err := json.Unmarshal(body, req)
	helper.PanicIfError(err)

	res := controller.UserService.Register(c, *req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *UserControllerImpl) Update(c *fiber.Ctx) error {
	body := c.Body()
	req := new(web.UserUpdateRequest)
	err := json.Unmarshal(body, req)
	helper.PanicIfError(err)

	user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    req.Username = claims["username"].(string)

	res := controller.UserService.Update(c, *req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *UserControllerImpl) Login(c *fiber.Ctx) error {
	body := c.Body()
	req := new(web.LoginRequest)
	err := json.Unmarshal(body, req)
	helper.PanicIfError(err)

	res := controller.UserService.Login(c, *req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}
