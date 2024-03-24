package helper

import (
	"slices"
	"todo-api/model/domain"
)

type Payload struct {
	Pupdate []domain.TodoCategory
	Pdelete []domain.TodoCategory
}

type ValidateResult struct {
	Categories []domain.TodoCategory
	Ids        []string
}

func ValidateCategories(reqCategories []domain.TodoCategory, categories []domain.Category) (ValidateResult, bool) {
	var invalidCategories []string
	var cat []domain.TodoCategory
	if len(categories) == 0 {
		for _, v := range reqCategories {
			invalidCategories = append(invalidCategories, v.CategoryId)
		}
		return ValidateResult{Ids: invalidCategories, Categories: reqCategories}, false
	}
	var same []string
	for _, req := range reqCategories {
		for _, c := range categories {
			if req.CategoryId == c.Id {
				same = append(same, req.CategoryId)
				cat = append(cat, domain.TodoCategory{TodoId: req.TodoId, CategoryId: c.Id, Name: c.Name})
				break
			}
		}
	}
	if len(same) == len(reqCategories) {
		return ValidateResult{Categories: cat, Ids: invalidCategories}, true
	}
	for _, req := range reqCategories {
		ok := slices.Contains(same, req.CategoryId)
		if !ok {
			invalidCategories = append(invalidCategories, req.CategoryId)
		}
	}
	if len(invalidCategories) > 0 {
		return ValidateResult{Ids: invalidCategories, Categories: cat}, false
	}
	return ValidateResult{}, true
}

func ValidateCategoriesForUpdate(reqC []domain.TodoCategory, dbC []domain.TodoCategory) Payload {
	if len(reqC) == 0 {
		return Payload{}
	}

	var same []string
	for _, v := range reqC {
		for _, v2 := range dbC {
			if v.CategoryId == v2.CategoryId {
				same = append(same, v.CategoryId)
				break
			}
		}
	}
	var payloadUpdate []domain.TodoCategory
	var payloadDelete []domain.TodoCategory
	for _, req := range reqC {
		ok := slices.Contains(same, req.CategoryId)
		if !ok {
			payloadUpdate = append(payloadUpdate, domain.TodoCategory{
				CategoryId: req.CategoryId,
				TodoId:     req.TodoId,
				Name:       MatchCategoryName(dbC, req.CategoryId),
			})
		}
	}
	for _, db := range dbC {
		if db.TodoId == "" {
			continue
		}
		ok := slices.Contains(same, db.CategoryId)
		if !ok {
			payloadDelete = append(payloadDelete, domain.TodoCategory{
				CategoryId: db.CategoryId,
				TodoId:     db.TodoId,
				Name:       db.Name,
			})
		}
	}
	return Payload{
		Pupdate: payloadUpdate,
		Pdelete: payloadDelete,
	}
}
