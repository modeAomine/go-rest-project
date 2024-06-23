package repo

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
}

func NewRepo() *Repo {
	return &Repo{}
}
