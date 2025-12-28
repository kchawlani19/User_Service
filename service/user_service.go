package service

import (
	"errors"
	"go-basic-user-service/model"
	"go-basic-user-service/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user model.User) error {
	exists, err := s.repo.Exists(user.Id)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exists")
	}
	return s.repo.Save(user)
}

func (s *UserService) Get(id int) (model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(id int, name string) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	user.Name = name
	return s.repo.Update(user)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}
