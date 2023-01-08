package service

import (
	"crypto/sha1"
	"fmt"
	todo "todo/study"
	"todo/study/package/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const salt = "timofey"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user todo.RegUser) (int, error) {
	user.Password = a.GeneratePasswordHash(user.Password)
	id, err := a.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthService) LoginUser(user todo.User) (string, error) {
	user.Password = a.GeneratePasswordHash(user.Password)

	userId, err := a.repo.IsUserExist(user)
	if err != nil {
		return "", err
	}
	if userId == 0 {
		logrus.Error("Пользователь с такими параметрами не найден")
		return "", nil
	}

	token, err := a.GenerateJWT(userId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type loginJwt struct {
	jwt.StandardClaims
	Id int `json:"id"`
}

func (s *AuthService) GenerateJWT(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &loginJwt{
		Id: id,
	})

	return token.SignedString([]byte("StrongKey"))
}
