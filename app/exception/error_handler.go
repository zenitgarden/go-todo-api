package exception

import (
	"fmt"
	"strings"
	"todo-api/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	fmt.Println(err)
	if res, ok := validationError(c, err); ok {
		return c.JSON(res)
	}
	if res, ok := notFoundError(c, err); ok {
		return c.JSON(res)
	}
	if res, ok := uniqueError(c, err); ok {
		return c.JSON(res)
	}
	if res, ok := loginError(c, err); ok {
		return c.JSON(res)
	}

	c.Status(fiber.StatusInternalServerError)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
	})
}

func NotFoundHandler(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusNotFound,
		Status: "NOT_FOUND",
	})
}

func JWTerrorHandler(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusUnauthorized)
	return c.JSON(web.WebResponse{
		Code:   fiber.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data: fiber.Map{
			"message": "invalid or expired token",
		},
	})
}

func validationError(c *fiber.Ctx, err any) (web.WebResponse, bool) {
	exception, ok := err.(validator.ValidationErrors)
	msg := map[string]string{}
	for _, err := range exception {
		field := strings.ToLower(err.Field())
		switch err.Tag() {
		case "required":
			msg[field] = fmt.Sprintf("%s is required", field)
		case "min":
			msg[field] = fmt.Sprintf("%s value must be more than 1 character", field)
		case "max":
			msg[field] = fmt.Sprintf("%s value must be less than 200 characters", field)
		}
	}

	if ok {
		c.Status(fiber.StatusBadRequest)
		return web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   msg,
		}, true
	}
	return web.WebResponse{}, false
}

func notFoundError(c *fiber.Ctx, err any) (web.WebResponse, bool) {
	exception, ok := err.(NotFoundError)
	if ok {
		c.Status(fiber.StatusNotFound)
		tags, _ := exception.Tags.([]any)
		if len(tags) > 0 {
			var errorRes []fiber.Map
			tag := tags[0].([]string)
			for _, t := range tag {
				errorRes = append(errorRes, fiber.Map{
					"id":      t,
					"message": exception.Msg,
				})
			}
			return web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "NOT_FOUND",
				Data:   errorRes,
			}, true
		}
		return web.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: "NOT_FOUND",
			Data:   exception.Msg,
		}, true
	}
	return web.WebResponse{}, false
}

func uniqueError(c *fiber.Ctx, err any) (web.WebResponse, bool) {
	exception, ok := err.(UniqueError)
	if ok {
		c.Status(fiber.StatusBadRequest)
		return web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data: fiber.Map{
				exception.Field: exception.Msg,
			},
		}, true
	}
	return web.WebResponse{}, false
}

func loginError(c *fiber.Ctx, err any) (web.WebResponse, bool) {
	exception, ok := err.(LoginError)
	if ok {
		c.Status(fiber.StatusBadRequest)
		return web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   exception.Msg,
		}, true
	}
	return web.WebResponse{}, false
}
