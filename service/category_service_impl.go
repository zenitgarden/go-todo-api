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

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(c *fiber.Ctx, r web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, errorDB := service.DB.Begin()
	helper.PanicIfError(errorDB)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{Id: "category-" + uuid.NewString(), Name: r.Name, Owner: r.Owner}

	category = service.CategoryRepository.Save(c, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(c *fiber.Ctx, r web.CategoryUpdateRequest) web.CategoryResponse {
	tx, errorDB := service.DB.Begin()
	helper.PanicIfError(errorDB)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{Id: r.Id, Name: r.Name, Owner: r.Owner}
	dbCategory, err := service.CategoryRepository.FindById(c, tx, category.Id)

	if err != nil || dbCategory.Owner != r.Owner {
		panic(exception.NewNotFoundError("category is not found"))
	}

	err = service.Validate.Struct(r)
	helper.PanicIfError(err)

	service.CategoryRepository.Update(c, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(c *fiber.Ctx, r web.CategoryRequest) web.CategoryResponse {
	tx, errorDB := service.DB.Begin()
	helper.PanicIfError(errorDB)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{Id: r.Id, Owner: r.Owner}
	dbCategory, err := service.CategoryRepository.FindById(c, tx, category.Id)
	if err != nil || dbCategory.Owner != r.Owner {
		panic(exception.NewNotFoundError("category is not found"))
	}

	service.CategoryRepository.Delete(c, tx, category)
	return helper.ToCategoryResponse(dbCategory)
}

func (service *CategoryServiceImpl) FindAll(c *fiber.Ctx, owner string) []web.CategoryResponse {
	tx, errorDB := service.DB.Begin()
	helper.PanicIfError(errorDB)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(c, tx, owner)
	return helper.ToCategoriesResponse(categories)
}

func (service *CategoryServiceImpl) FindById(c *fiber.Ctx, r web.CategoryRequest) web.CategoryResponse {
	tx, errorDB := service.DB.Begin()
	helper.PanicIfError(errorDB)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{Id: r.Id, Owner: r.Owner}
	dbCategory, err := service.CategoryRepository.FindById(c, tx, category.Id)
	if err != nil || dbCategory.Owner != r.Owner {
		panic(exception.NewNotFoundError("category is not found"))
	}
	return helper.ToCategoryResponse(dbCategory)
}
