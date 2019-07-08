package mocks

import (
	"github.com/nandaryanizar/golang-webservice-example/entities"
	"github.com/stretchr/testify/mock"
)

// UserServiceMock struct
type UserServiceMock struct {
	mock.Mock
}

// AuthenticateUser mock method
func (m *UserServiceMock) AuthenticateUser(email, password string) (entities.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(entities.User), args.Error(1)
}

// FindUserByID mock method
func (m *UserServiceMock) FindUserByID(id int) (entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(entities.User), args.Error(1)
}