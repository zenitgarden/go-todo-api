package helper

import (
	"todo-api/model/domain"
)

func TodoIds(todos []domain.Todo) []string {
	var result []string
	for _, todo := range todos {
		result = append(result, todo.Id)
	}
	return result
}
