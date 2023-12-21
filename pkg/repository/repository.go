package repository

// DB handlers

type Auth interface {
}

type TodoList interface {
}

type Item interface {
}

type Repository struct {
	Auth
	TodoList
	Item
}

func NewRepository() *Repository {
	return &Repository{}
}
