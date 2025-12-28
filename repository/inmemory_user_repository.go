package repository

import (
	"errors"
	"go-basic-user-service/model"
	"sync"
)

type InMemoryUserRepository struct {
	users []model.User
	mu    sync.Mutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []model.User{},
	}
}

func (r *InMemoryUserRepository) Save(user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users = append(r.users, user)
	return nil
}

func (r *InMemoryUserRepository) GetByID(id int) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Id == id {
			return u, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func (r *InMemoryUserRepository) Update(user model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, u := range r.users {
		if u.Id == user.Id {
			r.users[i] = user
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *InMemoryUserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, u := range r.users {
		if u.Id == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *InMemoryUserRepository) Exists(id int) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Id == id {
			return true, nil
		}
	}
	return false, nil
}
