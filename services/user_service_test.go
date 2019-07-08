package services_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/nandaryanizar/golang-webservice-example/entities"
	"github.com/nandaryanizar/golang-webservice-example/repositories/mocks"
	"github.com/nandaryanizar/golang-webservice-example/services"
)

func TestAuthenticateUser(t *testing.T) {
	userRepo := new(mocks.UserRepositoryMock)
	userRepo.On("FindByEmail", "test.valid@gmail.com").Return(entities.User{
		ID:       1,
		Email:    "test.valid@gmail.com",
		Password: "$2a$10$FJjWKtgaJkZPEe/UHJEpg./YsQLBOjT./mA969guhJi.yrA3J9zl.",
	}, nil)

	userRepo.On("FindByEmail", "test.invalid@gmail.com").Return(entities.User{
		ID:       0,
		Email:    "test.invalid@gmail.com",
		Password: "$2a$10$FJjWKtgaJkZPEe/UHJEpg./YsQLBOjT./mA969guhJi.yrA3J9zl.",
	}, nil)

	userRepo.On("FindByEmail", "test.invalid@gmail.com").Return(entities.User{
		Email: "test.valid@gmail.com",
	}, errors.New("Email and/or password cannot be empty string"))

	userService := services.NewUserService(userRepo)

	cases := []struct {
		email    string
		password string
		result   entities.User
	}{
		{"test.valid@gmail.com", "testpassword", entities.User{ID: 1, Email: "test.valid@gmail.com", Password: ""}},
		{"test.invalid@gmail.com", "testpassword", entities.User{ID: 0, Email: "test.invalid@gmail.com", Password: ""}},
	}

	for _, tc := range cases {
		user, err := userService.AuthenticateUser(tc.email, tc.password)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(user, tc.result) {
			t.Errorf("Error: user authenticated expected %v, found %v", tc.result, user)
		}
	}
}
