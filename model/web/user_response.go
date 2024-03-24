package web

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
