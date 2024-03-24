package repository

import (
	"database/sql"
	"todo-api/model/domain"

	"github.com/gofiber/fiber/v2"
)

type CategoryRepository interface {
	Save(c *fiber.Ctx, tx *sql.Tx, category domain.Category) domain.Category
	Update(c *fiber.Ctx, tx *sql.Tx, category domain.Category) domain.Category
	Delete(c *fiber.Ctx, tx *sql.Tx, category domain.Category)
	FindById(c *fiber.Ctx, tx *sql.Tx, categoryId string) (domain.Category, error)
	FindAll(c *fiber.Ctx, tx *sql.Tx, owner string) []domain.Category
	SaveTodoCategories(c *fiber.Ctx, tx *sql.Tx, categories []domain.TodoCategory)
	FindTodoCategoriesByTodoId(c *fiber.Ctx, tx *sql.Tx, todoId string) []domain.TodoCategory
	FindTodoCategoriesByTodoIds(c *fiber.Ctx, tx *sql.Tx, todoIds []string) []domain.TodoCategory
	FindAllById(c *fiber.Ctx, tx *sql.Tx, categories []domain.TodoCategory, owner string) []domain.Category
	FindAllForUpdate(c *fiber.Ctx, tx *sql.Tx, categories []domain.TodoCategory, owner string) []domain.TodoCategory
	DeleteTodoCategories(c *fiber.Ctx, tx *sql.Tx, categories []domain.TodoCategory)
}
