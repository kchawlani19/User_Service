package repository

import (
	"go-basic-user-service/model"
)

type UserRepository interface {
	Save(user model.User) error
	GetByID(id int) (model.User, error)
	Update(user model.User) error
	Delete(id int) error
	Exists(id int) (bool, error)
}
