package repository

import (
	"database/sql"
	"todo-api/model/domain"
	"todo-api/model/web"

	"github.com/gofiber/fiber/v2"
)

type TodoRepository interface {
	Save(c *fiber.Ctx, tx *sql.Tx, todo domain.Todo) domain.Todo
	Update(c *fiber.Ctx, tx *sql.Tx, todo domain.Todo) domain.Todo
	Delete(c *fiber.Ctx, tx *sql.Tx, todo domain.Todo)
	FindById(c *fiber.Ctx, tx *sql.Tx, todoId string) (domain.Todo, error)
	FindAll(c *fiber.Ctx, tx *sql.Tx, owner string) []domain.Todo
	Search(c *fiber.Ctx, tx *sql.Tx, owner string, qs web.TQueryString) []domain.Todo
}
