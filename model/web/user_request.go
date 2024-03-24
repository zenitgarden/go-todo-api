package web

type UserCreateRequest struct {
	Name     string `validate:"required,max=200,min=1" json:"name"`
	Username string `validate:"required,max=200,min=1" json:"username"`
	Password string `validate:"required,max=200,min=1" json:"password"`
}

type UserUpdateRequest struct {
	Username string `validate:"required,max=200,min=1" json:"username"`
	Name     string `validate:"required,max=200,min=1" json:"name"`
}

type LoginRequest struct {
	Username string `validate:"required,max=200,min=1" json:"username"`
	Password string `validate:"required,max=200,min=1" json:"password"`
}
