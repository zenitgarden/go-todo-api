package web

type CategoryCreateRequest struct {
	Name  string `validate:"required,max=200,min=1" json:"name"`
	Owner string `validate:"required"`
}

type CategoryUpdateRequest struct {
	Id    string `validate:"required"`
	Name  string `validate:"required,max=200,min=1" json:"name"`
	Owner string `validate:"required"`
}

type CategoryRequest struct {
	Id    string `validate:"required"`
	Owner string `validate:"required"`
}

type TodoCategoriesRequest struct {
	Id    string `validate:"required"`
	Name  string `validate:"required,max=200,min=1" json:"name"`
}
