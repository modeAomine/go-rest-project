package service

import "web-go/pkg/repo"

type Auth interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Auth
	TodoList
	TodoItem
}

func NewService(repos *repo.Repo) *Service {
	return &Service{}
}
