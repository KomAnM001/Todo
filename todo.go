package storage

import (
	"atodo/models"
)

type Todo interface {
	CreateUser(in *models.User) (resp *models.User, err error)
	GetByID(in *models.User) (resp *models.User, err error)
	GetAll() (resp []models.User, err error)
	UpdateUser(in *models.User) (resp *models.User, err error)
	DeleteUser(in *models.User) (err error)
}

type StoreI interface {
	Todo() Todo
}
