package service

import (
	"crypto/sha1"
	"fmt"
	tasks_app "tasks-list-project"
	"tasks-list-project/pkg/repository"
)

const salt = "sadfgyu12345ghjdsgjhkerthjh574"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user tasks_app.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
