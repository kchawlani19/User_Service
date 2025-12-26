package repository

import (
	"go-basic-user-service/model"
)

type UserRepository struct {
	users []model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []model.User{},
	}
}

func (r *UserRepository) Save(user model.User) {
	r.users = append(r.users, user)
}

func (r *UserRepository) ExistsByID(id int) bool {
	for i := 0; i < len(r.users); i++ {
		if r.users[i].Id == id {
			return true
		}
	}
	return false
}

func (r *UserRepository) GetById(id int) (model.User, bool) {
	for _, u := range r.users {
		if u.Id == id {
			return u, true
		}
	}
	return model.User{}, false
}

func (r *UserRepository) Update(user model.User) {
	for i := 0; i < len(r.users); i++ {
		if r.users[i].Id == user.Id {
			r.users[i] = user
			return
		}
	}
}

func (r *UserRepository) Delete(id int) {
	for i := 0; i < len(r.users); i++ {
		if r.users[i].Id == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return
		}
	}
}
