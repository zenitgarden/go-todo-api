package repository

import (
	"database/sql"
	"todo-api/model/domain"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Save(c *fiber.Ctx, tx *sql.Tx, user domain.User) domain.User
	Update(c *fiber.Ctx, tx *sql.Tx, user domain.User) domain.User
	FindByUsername(c *fiber.Ctx, tx *sql.Tx, username string) (domain.User, error)
}
