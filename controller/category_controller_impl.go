package controller

import (
	"encoding/json"
	"todo-api/helper"
	"todo-api/model/web"
	"todo-api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Save(c *fiber.Ctx) error {
	body := c.Body()
	req := new(web.CategoryCreateRequest)
	err := json.Unmarshal(body, req)
	helper.PanicIfError(err)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.Owner = claims["username"].(string)

	res := controller.CategoryService.Create(c, *req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *CategoryControllerImpl) Update(c *fiber.Ctx) error {
	body := c.Body()
	req := new(web.CategoryUpdateRequest)
	err := json.Unmarshal(body, req)
	helper.PanicIfError(err)

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.Owner = claims["username"].(string)
	req.Id = c.Params("categoryId")

	res := controller.CategoryService.Update(c, *req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *CategoryControllerImpl) Delete(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req := web.CategoryRequest{
		Id:    c.Params("categoryId"),
		Owner: claims["username"].(string),
	}
	res := controller.CategoryService.Delete(c, req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *CategoryControllerImpl) FindAll(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	owner := claims["username"].(string)

	res := controller.CategoryService.FindAll(c, owner)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

func (controller *CategoryControllerImpl) FindById(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req := web.CategoryRequest{
		Id:    c.Params("categoryId"),
		Owner: claims["username"].(string),
	}

	res := controller.CategoryService.FindById(c, req)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   res,
	})
}
