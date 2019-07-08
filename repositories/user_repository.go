package repositories

import (
	"github.com/nandaryanizar/fury"
	"github.com/nandaryanizar/golang-webservice-example/entities"
)

// UserRepository interface
type UserRepository interface {
	FindByEmail(email string) (entities.User, error)
	FindByID(id int) (entities.User, error)
}

type userRepository struct {
	db *fury.DB
}

// NewUserRepository factory function
func NewUserRepository(db *fury.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// FindByEmail method
func (ur *userRepository) FindByEmail(email string) (entities.User, error) {
	user := entities.User{
		Email: email,
	}
	if err := ur.db.First(&user, fury.Table("users"), fury.Where(fury.IsEqualsTo("email", user.Email))); err != nil {
		return user, err
	}

	return user, nil
}

// FindByID method
func (ur *userRepository) FindByID(id int) (entities.User, error) {
	user := entities.User{
		ID: id,
	}
	if err := ur.db.First(&user, fury.Table("users"), fury.Where(fury.IsEqualsTo("id", id))); err != nil {
		return user, err
	}

	return user, nil
}
