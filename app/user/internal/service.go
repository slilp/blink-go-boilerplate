package internal

import (
	user "blink-go-gin-boilerplate/app/user/api"
)


func NewService(repo user.Repository) user.Service {
	return &service{repo: repo}
}

type service struct {
	repo user.Repository
}

func (s *service) Create(user *user.UserEntity) (*user.UserEntity, error) {
	return s.repo.Create(user)
}


func (s *service) GetByUsername(username string) (*user.UserEntity, error) {
	return s.repo.GetByUsername(username)
}
