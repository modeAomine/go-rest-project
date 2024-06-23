package repo

import "github.com/jinzhu/gorm"

type Auth interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repo struct {
	Auth
	TodoList
	TodoItem
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}
