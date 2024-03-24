package repository

import (
	"database/sql"
	"errors"
	"todo-api/helper"
	"todo-api/model/domain"

	"github.com/gofiber/fiber/v2"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(c *fiber.Ctx, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(username,name,password) VALUES(?,?,?)"
	_, err := tx.ExecContext(c.Context(), SQL, user.Username, user.Name, user.Password)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) Update(c *fiber.Ctx, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET name = ? WHERE username = ? "
	_, err := tx.ExecContext(c.Context(), SQL, user.Name, user.Username)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) FindByUsername(c *fiber.Ctx, tx *sql.Tx, username string) (domain.User, error) {
	SQL := "SELECT username, password, name FROM users WHERE username = ?"
	rows, err := tx.QueryContext(c.Context(), SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Password, &user.Name)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}
