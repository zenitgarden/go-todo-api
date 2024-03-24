package domain

import "time"

type Todo struct {
	Id          string
	Title       string
	Description string
	Owner       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}