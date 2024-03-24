package controller

import (
	"encoding/json"
	"todo-api/helper"
	"todo-api/model/web"
	"todo-api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

func (controller *TodoControllerImpl) Save(c *fiber.Ctx) error {
	body := c.Body()
	req := new(web.TodoCreateRequest)
	err := json.Unmarshal(body, req)
	helper.PanicIfError(err)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.Owner = claims["username"].(string)

	res := controller.TodoService.Create(c, *req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *TodoControllerImpl) Update(c *fiber.Ctx) error {
	body := c.Body()
	req := new(web.TodoUpdateRequest)
	err := json.Unmarshal(body, req)
	helper.PanicIfError(err)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.Owner = claims["username"].(string)
	req.Id = c.Params("todoId")

	res := controller.TodoService.Update(c, *req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *TodoControllerImpl) Delete(c *fiber.Ctx) error {
	body := c.Body()
	req := new(web.TodoRequest)
	err := json.Unmarshal(body, req)
	helper.PanicIfError(err)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.Owner = claims["username"].(string)
	req.Id = c.Params("todoId")

	res := controller.TodoService.Delete(c, *req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *TodoControllerImpl) FindAll(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	owner := claims["username"].(string)

	m := c.Queries()
	qs := web.TQueryString{
		Search:   m["search"],
		Category: m["category"],
	}

	res := controller.TodoService.FindAll(c, owner, qs)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *TodoControllerImpl) FindById(c *fiber.Ctx) error {
	req := new(web.TodoRequest)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.Owner = claims["username"].(string)
	req.Id = c.Params("todoId")

	res := controller.TodoService.FindById(c, *req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}
