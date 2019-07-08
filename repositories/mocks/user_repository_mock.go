package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/nandaryanizar/golang-webservice-example/entities"
)

// UserRepositoryMock struct
type UserRepositoryMock struct {
	mock.Mock
}

// FindByEmail mock method
func (m *UserRepositoryMock) FindByEmail(email string) (entities.User, error) {
	args := m.Called(email)
	return args.Get(0).(entities.User), args.Error(1)
}

// FindByID mock method
func (m *UserRepositoryMock) FindByID(id int) (entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(entities.User), args.Error(1)
}
