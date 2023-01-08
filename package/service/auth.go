package service

import (
	"crypto/sha1"
	"fmt"
	todo "todo/study"
	"todo/study/package/repository"
)

const salt = "timofey"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = a.GeneratePasswordHash(user.Password)
	id, err := a.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *AuthService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
