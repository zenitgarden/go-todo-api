package web

import "time"

type TodoResponse struct {
	Id          string             `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Categories  []CategoryResponse `json:"categories"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type TodoDeleteResponse struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}
