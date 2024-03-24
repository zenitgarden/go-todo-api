package web

type TodoCreateRequest struct {
	Title       string                  `validate:"required,max=200,min=1" json:"title"`
	Description string                  `validate:"required,max=200,min=1" json:"description"`
	Owner       string                  `validate:"required" json:"owner"`
	Categories  []TodoCategoriesRequest `json:"categories"`
}

type TodoUpdateRequest struct {
	Id          string                  `validate:"required" json:"id"`
	Title       string                  `validate:"required,max=200,min=1" json:"title"`
	Description string                  `validate:"required,max=200,min=1" json:"description"`
	Owner       string                  `validate:"required" json:"owner"`
	Categories  []TodoCategoriesRequest `json:"categories"`
}

type TodoRequest struct {
	Id    string `validate:"required"`
	Owner string `validate:"required"`
}

type TQueryString struct{
	Search string
	Category string
}
