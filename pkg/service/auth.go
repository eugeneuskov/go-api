package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-api/models"
	"go-api/pkg/repository"
	"log"
	"time"
)

const (
	salt            = "fxng83tinEDD,%&84tyt@3hn8c!5ty93mHJp2wql,lqs;"
	tokenSigningKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL        = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repository.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	log.Printf("token: %v\n", token)

	return token.SignedString([]byte(tokenSigningKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
