package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"todo-api/helper"
	"todo-api/model/domain"

	"github.com/gofiber/fiber/v2"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(c *fiber.Ctx, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(id,name,owner) VALUES(?,?,?)"
	_, err := tx.ExecContext(c.Context(), SQL, category.Id, category.Name, category.Owner)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Update(c *fiber.Ctx, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(c.Context(), SQL, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(c *fiber.Ctx, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE from categories WHERE id = ?"
	_, err := tx.ExecContext(c.Context(), SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(c *fiber.Ctx, tx *sql.Tx, categoryId string) (domain.Category, error) {
	SQL := "SELECT id, name, owner FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(c.Context(), SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name, &category.Owner)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(c *fiber.Ctx, tx *sql.Tx, owner string) []domain.Category {
	SQL := "SELECT id, name FROM categories WHERE owner = ?"
	rows, err := tx.QueryContext(c.Context(), SQL, owner)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	err = rows.Err()
	helper.PanicIfError(err)

	return categories
}

func (repository *CategoryRepositoryImpl) SaveTodoCategories(c *fiber.Ctx, tx *sql.Tx, categories []domain.TodoCategory) {
	var questionMark string
	var values []interface{}
	for i, v := range categories {
		questionMark = questionMark + "(?,?),"
		if i == len(categories)-1 {
			questionMark = questionMark[:len(questionMark)-1]
		}
		values = append(values, v.TodoId, v.CategoryId)
	}
	SQL := fmt.Sprintf("INSERT INTO todo_category(todo_id, category_id) VALUES %s", questionMark)
	_, err := tx.ExecContext(c.Context(), SQL, values...)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindTodoCategoriesByTodoId(c *fiber.Ctx, tx *sql.Tx, todoId string) []domain.TodoCategory {
	SQL := "SELECT tc.todo_id, tc.category_id, c.name FROM todo_category as tc JOIN todos as t ON t.id =tc.todo_id JOIN categories as c ON c.id=tc.category_id WHERE tc.todo_id = ?;"
	rows, err := tx.QueryContext(c.Context(), SQL, todoId)
	helper.PanicIfError(err)

	var todoCategories []domain.TodoCategory
	for rows.Next() {
		todoCategory := domain.TodoCategory{}
		err := rows.Scan(&todoCategory.TodoId, &todoCategory.CategoryId, &todoCategory.Name)
		helper.PanicIfError(err)
		todoCategories = append(todoCategories, todoCategory)
	}
	return todoCategories
}

func (repository *CategoryRepositoryImpl) FindTodoCategoriesByTodoIds(c *fiber.Ctx, tx *sql.Tx, todoIds []string) []domain.TodoCategory {
	var questionMark string
	var values []interface{}
	for i, v := range todoIds {
		questionMark = questionMark + "?,"
		values = append(values, v)
		if i == len(todoIds)-1 {
			questionMark = questionMark[:len(questionMark)-1]
		}
	}
	SQL := fmt.Sprintf("SELECT tc.todo_id, tc.category_id, c.name FROM todo_category as tc JOIN todos as t ON t.id =tc.todo_id JOIN categories as c ON c.id=tc.category_id WHERE tc.todo_id IN (%s)", questionMark)
	rows, err := tx.QueryContext(c.Context(), SQL, values...)
	helper.PanicIfError(err)

	var todoCategories []domain.TodoCategory
	for rows.Next() {
		todoCategory := domain.TodoCategory{}
		err := rows.Scan(&todoCategory.TodoId, &todoCategory.CategoryId, &todoCategory.Name)
		helper.PanicIfError(err)
		todoCategories = append(todoCategories, todoCategory)
	}
	return todoCategories
}

func (repository *CategoryRepositoryImpl) FindAllById(c *fiber.Ctx, tx *sql.Tx, categories []domain.TodoCategory, owner string) []domain.Category {
	var questionMark string
	var values []interface{}
	for i, v := range categories {
		questionMark = questionMark + "?,"
		values = append(values, v.CategoryId)
		if i == len(categories)-1 {
			questionMark = questionMark[:len(questionMark)-1]
		}
	}
	SQL := fmt.Sprintf("SELECT id,name FROM categories WHERE id IN (%s) AND owner = ?", questionMark)
	values = append(values, owner)
	rows, err := tx.QueryContext(c.Context(), SQL, values...)
	helper.PanicIfError(err)
	defer rows.Close()

	var resultCategories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		resultCategories = append(resultCategories, category)
	}

	err = rows.Err()
	helper.PanicIfError(err)

	return resultCategories
}

func (repository *CategoryRepositoryImpl) FindAllForUpdate(c *fiber.Ctx, tx *sql.Tx, categories []domain.TodoCategory, owner string) []domain.TodoCategory {
	var questionMark string
	var values []interface{}
	for i, v := range categories {
		questionMark = questionMark + "?,"
		values = append(values, v.CategoryId)
		if i == len(categories)-1 {
			questionMark = questionMark[:len(questionMark)-1]
		}
	}
	SQL := fmt.Sprintf("SELECT * FROM (SELECT t1.id as c_id, t2.todo_id as t_id, t1.name FROM categories as t1 LEFT JOIN todo_category as t2 ON t1.id=t2.category_id WHERE id IN (%s) AND owner = ?) AS sub1 WHERE NOT EXISTS (SELECT * FROM (SELECT t1.id as c_id, t2.todo_id as t_id, t1.name FROM categories as t1 LEFT JOIN todo_category as t2 ON t1.id=t2.category_id WHERE id IN (%s) AND owner = ?) AS sub2 WHERE c_id = sub1.c_id AND t_id <> sub1.t_id  AND t_id = ?)", questionMark, questionMark)
	values = append(values, owner)
	values = append(values, values...)
	values = append(values, categories[0].TodoId)
	fmt.Println(values...)
	fmt.Println(SQL)
	rows, err := tx.QueryContext(c.Context(), SQL, values...)
	helper.PanicIfError(err)
	defer rows.Close()

	var todoCategories []domain.TodoCategorySqlString
	fmt.Println("Query success")
	for rows.Next() {
		todoCategory := domain.TodoCategorySqlString{}
		err := rows.Scan(&todoCategory.CategoryId, &todoCategory.TodoId, &todoCategory.Name)
		helper.PanicIfError(err)
		todoCategories = append(todoCategories, todoCategory)
	}

	err = rows.Err()
	helper.PanicIfError(err)

	return helper.ToCategoryStirng(todoCategories)
}

func (repository *CategoryRepositoryImpl) DeleteTodoCategories(c *fiber.Ctx, tx *sql.Tx, categories []domain.TodoCategory) {
	var questionMark string
	var values []interface{}
	for i, v := range categories {
		questionMark = questionMark + "?,"
		if i == len(categories)-1 {
			questionMark = questionMark[:len(questionMark)-1]
		}
		values = append(values, v.CategoryId)
	}
	SQL := fmt.Sprintf("DELETE from todo_category WHERE category_id IN (%s) AND todo_id = ?", questionMark)
	values = append(values, categories[0].TodoId)
	_, err := tx.ExecContext(c.Context(), SQL, values...)
	helper.PanicIfError(err)
}
