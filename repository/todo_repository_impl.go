package repository

import (
	"database/sql"
	"errors"
	"todo-api/helper"
	"todo-api/model/domain"
	"todo-api/model/web"

	"github.com/gofiber/fiber/v2"
)

type TodoRepositoryImpl struct{}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (repository *TodoRepositoryImpl) Save(c *fiber.Ctx, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "INSERT INTO todos(id, title, description, owner) VALUES (?,?,?,?)"
	_, err := tx.ExecContext(c.Context(), SQL, todo.Id, todo.Title, todo.Description, todo.Owner)
	helper.PanicIfError(err)
	return todo
}

func (repository *TodoRepositoryImpl) Update(c *fiber.Ctx, tx *sql.Tx, todo domain.Todo) domain.Todo {
	var SQL string
	var values []interface{}
	if todo.Description == "" {
		SQL = "UPDATE todos SET title = ? WHERE id = ? AND owner = ?"
		values = append(values, todo.Title, todo.Id, todo.Owner)
	} else {
		SQL = "UPDATE todos SET title = ?, description = ? WHERE id = ? AND owner = ?"
		values = append(values, todo.Title, todo.Description, todo.Id, todo.Owner)
	}
	_, err := tx.ExecContext(c.Context(), SQL, values...)
	helper.PanicIfError(err)
	return todo
}

func (repository *TodoRepositoryImpl) Delete(c *fiber.Ctx, tx *sql.Tx, todo domain.Todo) {
	SQL := "DELETE FROM todos WHERE id = ? AND owner = ?"
	_, err := tx.ExecContext(c.Context(), SQL, todo.Id, todo.Owner)
	helper.PanicIfError(err)
}

func (repository *TodoRepositoryImpl) FindAll(c *fiber.Ctx, tx *sql.Tx, owner string) []domain.Todo {
	SQL := "SELECT id,title,description,created_at, updated_at FROM todos WHERE owner = ?"
	rows, err := tx.QueryContext(c.Context(), SQL, owner)
	helper.PanicIfError(err)

	var todos []domain.Todo
	for rows.Next() {
		todo := domain.Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		helper.PanicIfError(err)
		todos = append(todos, todo)
	}

	err = rows.Err()
	helper.PanicIfError(err)
	return todos
}

func (repository *TodoRepositoryImpl) FindById(c *fiber.Ctx, tx *sql.Tx, todoId string) (domain.Todo, error) {
	SQL := "SELECT id,title,owner,description,created_at, updated_at FROM todos WHERE id = ?"
	rows, err := tx.QueryContext(c.Context(), SQL, todoId)
	helper.PanicIfError(err)
	defer rows.Close()

	todo := domain.Todo{}
	if rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Owner, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		helper.PanicIfError(err)
		return todo, nil
	} else {
		return todo, errors.New("todo is not found")
	}
}

func (repository *TodoRepositoryImpl) Search(c *fiber.Ctx, tx *sql.Tx, owner string, qs web.TQueryString) []domain.Todo {
	SQL := "SELECT t.id, t.title, t.description, t.created_at, t.updated_at FROM todo_category as tc JOIN todos as t ON t.id =tc.todo_id JOIN categories as c ON c.id=tc.category_id WHERE"
	var values []interface{}
	count := 0
	if qs.Category != "" {
		SQL += " c.name = ?"
		count++
		values = append(values, qs.Category)
	}
	if qs.Search != "" {
		if count == 1 {
			SQL += " AND"
		}
		SQL += " t.title LIKE ? OR t.description LIKE ?"
		v := "%" + qs.Search + "%"
		values = append(values, v, v)
	}
	rows, err := tx.QueryContext(c.Context(), SQL, values...)
	helper.PanicIfError(err)

	var todos []domain.Todo
	for rows.Next() {
		todo := domain.Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		helper.PanicIfError(err)
		todos = append(todos, todo)
	}

	err = rows.Err()
	helper.PanicIfError(err)
	return todos
}
