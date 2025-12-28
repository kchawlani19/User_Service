package repository

import (
	"go-basic-user-service/model"
	"sync"
)

type UserRepository interface {
	ExistsByID(id int) bool
	Save(user model.User)
	GetById(id int) (model.User, bool)
	Update(user model.User)
	Delete(id int)
}

type InMemoryUserRepository struct {
	users []model.User
	mu    sync.Mutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []model.User{},
	}
}

func (r *InMemoryUserRepository) Save(user model.User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users = append(r.users, user)
}

func (r *InMemoryUserRepository) ExistsByID(id int) bool {
	for i := 0; i < len(r.users); i++ {
		if r.users[i].Id == id {
			return true
		}
	}
	return false
}

func (r *InMemoryUserRepository) GetById(id int) (model.User, bool) {
	for _, u := range r.users {
		if u.Id == id {
			return u, true
		}
	}
	return model.User{}, false
}

func (r *InMemoryUserRepository) Update(user model.User) {
	for i := 0; i < len(r.users); i++ {
		if r.users[i].Id == user.Id {
			r.users[i] = user
			return
		}
	}
}

func (r *InMemoryUserRepository) Delete(id int) {
	for i := 0; i < len(r.users); i++ {
		if r.users[i].Id == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return
		}
	}
}
