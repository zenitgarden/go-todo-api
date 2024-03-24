package domain

import "database/sql"

type Category struct {
	Id    string
	Name  string
	Owner string
}

type TodoCategory struct {
	TodoId     string
	CategoryId string
	Name       string
}

type TodoCategorySqlString struct {
	TodoId     sql.NullString
	CategoryId sql.NullString
	Name       sql.NullString
}
