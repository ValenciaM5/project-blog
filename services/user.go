package services

import (
	"errors"
	"project-blog/models"
	"project-blog/repository"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type UserService interface {
	Register(user *models.User) error
	Login(email, password string) (string, error)
	GetUser(id string) (*models.User, error)
}

type userService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{Repo: repo}
}

func (s *userService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.Repo.Create(user)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *userService) GetUser(id string) (*models.User, error) {
	return s.Repo.FindByID(id)
}