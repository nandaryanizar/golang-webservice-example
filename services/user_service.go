package services

import (
	"errors"

	"github.com/nandaryanizar/golang-webservice-example/entities"
	"github.com/nandaryanizar/golang-webservice-example/repositories"
	"golang.org/x/crypto/bcrypt"
)

// UserService interface
type UserService interface {
	AuthenticateUser(email string, password string) (entities.User, error)
	FindUserByID(id int) (entities.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

// NewUserService create UserService
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) FindUserByID(id int) (entities.User, error) {
	var user entities.User

	if id <= 0 {
		return user, errors.New("Invalid user ID")
	}

	user, err := us.repo.FindByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (us *userService) AuthenticateUser(email string, password string) (entities.User, error) {
	user := entities.User{}

	if email == "" || password == "" {
		return user, errors.New("Email and/or password cannot be empty string")
	}

	user, err := us.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if !us.isHashAndPasswordSame(user.Password, password) {
		return user, errors.New("Invalid username and/or password")
	}
	user.Password = ""

	return user, nil
}

func (us *userService) isHashAndPasswordSame(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
