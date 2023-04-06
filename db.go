package todo

import (
	"atodo/storage"
	"fmt"
	"os"
)

type Store struct {
	db   *os.File
	todo storage.Todo
}

func NewStore() (storage.StoreI, error) {

	file, err := os.OpenFile("/home/komilov/my_projects/atodo/todo.json", os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("error while open file")
	}

	return &Store{
		db:   file,
		todo: NewTodo(file),
	}, nil
}

func (s *Store) Todo() storage.Todo {
	if s.todo == nil {
		return NewTodo(s.db)
	}
	return s.todo
}
