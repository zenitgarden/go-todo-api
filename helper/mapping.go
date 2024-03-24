package helper

import (
	"todo-api/model/domain"
	"todo-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoriesResponse(categories []domain.Category) []web.CategoryResponse {
	var categoryResponse []web.CategoryResponse

	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}
	return categoryResponse
}

func ToTodoCategories(todoId string, categories []web.TodoCategoriesRequest) []domain.TodoCategory {
	var TodoCategory []domain.TodoCategory

	for _, category := range categories {
		TodoCategory = append(TodoCategory, domain.TodoCategory{
			TodoId:     string(todoId),
			CategoryId: category.Id,
		})
	}
	return TodoCategory
}

func ToSliceCategories(todoCategories []domain.TodoCategory) []domain.Category {
	var categories []domain.Category
	for _, category := range todoCategories {
		categories = append(categories, domain.Category{Id: category.CategoryId, Name: category.Name})
	}
	return categories
}

func ToTodoResponse(todo domain.Todo, categories []domain.Category) web.TodoResponse {
	return web.TodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		Categories:  ToCategoriesResponse(categories),
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Username: user.Username,
		Name:     user.Name,
	}
}

func ToCategoryFromTodoCategory(todoCategories []domain.TodoCategory) []domain.Category {
	var result []domain.Category
	for _, v := range todoCategories {
		result = append(result, domain.Category{
			Id:   v.CategoryId,
			Name: v.Name,
		})
	}
	return result
}

func ToCategoryStirng(todoCategories []domain.TodoCategorySqlString) []domain.TodoCategory {
	var result []domain.TodoCategory

	for _, v := range todoCategories {
		result = append(result, domain.TodoCategory{
			TodoId:     v.TodoId.String,
			CategoryId: v.CategoryId.String,
			Name:       v.Name.String,
		})
	}
	return result
}

func MatchCategoryName(c []domain.TodoCategory, id string) string {
	var result string
	for _, v := range c {
		if v.CategoryId == id {
			result = v.Name
			break
		}
	}
	return result
}

func MappingTodo(todos []domain.Todo, categories []domain.TodoCategory) []web.TodoResponse {
	var result []web.TodoResponse
	for _, t := range todos {
		var tempCategories []web.CategoryResponse
		for _, c := range categories {
			if t.Id == c.TodoId {
				tempCategories = append(tempCategories, web.CategoryResponse{
					Id:   c.CategoryId,
					Name: c.Name,
				})
			}
		}
		result = append(result, web.TodoResponse{
			Id:          t.Id,
			Title:       t.Title,
			Description: t.Description,
			Categories:  tempCategories,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		})
	}
	return result
}
