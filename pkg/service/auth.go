package service

import (
	"crypto/sha1"
	"fmt"
	"go-api/models"
	"go-api/pkg/repository"
)

const salt = "fxng83tinEDD,%&84tyt@3hn8c!5ty93mHJp2wql,lqs;"

type AuthService struct {
	repository repository.Authorization
}

func NewAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) CreateUser(user *models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repository.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
