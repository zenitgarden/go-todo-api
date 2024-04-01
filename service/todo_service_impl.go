package service

import (
	"database/sql"
	"todo-api/app/exception"
	"todo-api/helper"
	"todo-api/model/domain"
	"todo-api/model/web"
	"todo-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoServiceImpl struct {
	TodoRepository     repository.TodoRepository
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewTodoService(todoRepository repository.TodoRepository, categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository:     todoRepository,
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *TodoServiceImpl) Create(c *fiber.Ctx, r web.TodoCreateRequest) web.TodoResponse {
	err := service.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo := domain.Todo{Id: "todo-" + uuid.NewString(), Title: r.Title, Description: r.Description, Owner: r.Owner}
	todo = service.TodoRepository.Save(c, tx, todo)
	var categories []domain.Category
	if len(r.Categories) > 0 {
		todoCategories := helper.ToTodoCategories(todo.Id, r.Categories)
		categories = service.CategoryRepository.FindAllById(c, tx, todoCategories, todo.Owner)
		values, valid := helper.ValidateCategories(todoCategories, categories)
		if valid {
			service.CategoryRepository.SaveTodoCategories(c, tx, todoCategories)
		} else {
			panic(exception.NewNotFoundError("category is not found", values.Ids))
		}
	}
	todo, _ = service.TodoRepository.FindById(c, tx, todo.Id)
	return helper.ToTodoResponse(todo, categories)
}

func (service *TodoServiceImpl) Update(c *fiber.Ctx, r web.TodoUpdateRequest) web.TodoResponse {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(c, tx, r.Id)
	if err != nil || todo.Owner != r.Owner {
		panic(exception.NewNotFoundError("todo is not found"))
	}
	err = service.Validate.Struct(r)
	helper.PanicIfError(err)

	todo.Title = r.Title
	todo.Description = r.Description
	todo.Owner = r.Owner

	todo = service.TodoRepository.Update(c, tx, todo)
	var categories []domain.Category
	if len(r.Categories) > 0 {
		reqTodoCategories := helper.ToTodoCategories(todo.Id, r.Categories)
		categories = service.CategoryRepository.FindAllById(c, tx, reqTodoCategories, todo.Owner)
		result, valid := helper.ValidateCategories(reqTodoCategories, categories)
		reqTodoCategories = result.Categories
		if valid {
			todoCategoriesDB := service.CategoryRepository.FindTodoCategoriesByTodoId(c, tx, todo.Id)
			v := helper.ValidateCategoriesForUpdate(reqTodoCategories, todoCategoriesDB)
			if len(v.Pupdate) > 0 {
				service.CategoryRepository.SaveTodoCategories(c, tx, v.Pupdate)
			}
			if len(v.Pdelete) > 0 {
				service.CategoryRepository.DeleteTodoCategories(c, tx, v.Pdelete)
			}
		} else {
			panic(exception.NewNotFoundError("category is not found", result.Ids))
		}
	} else {
		todoCategoriesDB := service.CategoryRepository.FindTodoCategoriesByTodoId(c, tx, todo.Id)
		if len(todoCategoriesDB) > 0 {
			service.CategoryRepository.DeleteTodoCategories(c, tx, todoCategoriesDB)
		}
	}
	todo, _ = service.TodoRepository.FindById(c, tx, todo.Id)
	return helper.ToTodoResponse(todo, categories)
}

func (service *TodoServiceImpl) Delete(c *fiber.Ctx, r web.TodoRequest) web.TodoDeleteResponse {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)
	todo, err := service.TodoRepository.FindById(c, tx, r.Id)
	if err != nil || todo.Owner != r.Owner {
		panic(exception.NewNotFoundError("todo is not found"))
	}
	service.TodoRepository.Delete(c, tx, todo)
	return web.TodoDeleteResponse{
		Id:      todo.Id,
		Message: "todo has been deleted",
	}
}

func (service *TodoServiceImpl) FindAll(c *fiber.Ctx, owner string, qs web.TQueryString) []web.TodoResponse {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)
	if qs.Search != "" || qs.Category != "" {
		todos := service.TodoRepository.Search(c, tx, owner, qs)
		if len(todos) > 0 {
			categories := service.CategoryRepository.FindTodoCategoriesByTodoIds(c, tx, helper.TodoIds(todos))
			return helper.MappingTodo(todos, categories)
		}
	} else {
		todos := service.TodoRepository.FindAll(c, tx, owner)
		if len(todos) > 0 {
			categories := service.CategoryRepository.FindTodoCategoriesByTodoIds(c, tx, helper.TodoIds(todos))
			return helper.MappingTodo(todos, categories)
		}
	}
	return []web.TodoResponse{}
}

func (service *TodoServiceImpl) FindById(c *fiber.Ctx, r web.TodoRequest) web.TodoResponse {
	tx, errTx := service.DB.Begin()
	helper.PanicIfError(errTx)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(c, tx, r.Id)
	if err != nil || todo.Owner != r.Owner {
		panic(exception.NewNotFoundError("todo is not found"))
	}
	todoCategoriesDB := service.CategoryRepository.FindTodoCategoriesByTodoId(c, tx, todo.Id)
	return helper.ToTodoResponse(todo, helper.ToCategoryFromTodoCategory(todoCategoriesDB))
}
