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
	if s.repo.ExistsByID(user.Id) {
		return errors.New("duplicate id")
	}
	if user.Name == "" {
		return errors.New("name empty")
	}
	s.repo.Save(user)
	return nil
}

func (s *UserService) Get(id int) (model.User, error) {
	user, ok := s.repo.GetById(id)
	if !ok {
		return model.User{}, errors.New("not found")

	}
	return user, nil
}

func (s *UserService) Update(id int, name string) error {
	user, ok := s.repo.GetById(id)
	if !ok {
		return errors.New("not found")
	}

	user.Name = name
	s.repo.Update(user)
	return nil
}

func (s *UserService) Delete(id int) error {
	if !s.repo.ExistsByID(id) {
		return errors.New("not found")
	}

	s.repo.Delete(id)
	return nil
}
