package repositories_test

import (
	"testing"

	"github.com/nandaryanizar/fury"

	"github.com/nandaryanizar/golang-webservice-example/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nandaryanizar/golang-webservice-example/entities"
)

func TestFindByEmail(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Error opening database mock")
	}
	defer mockDB.Close()

	db, _ := fury.ConnectMock(mockDB)

	userRepo := repositories.NewUserRepository(db)

	mockUser := entities.User{
		ID:       1,
		Email:    "testvalid@gmail.com",
		Password: "testpassword123",
	}

	rows := sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(
		mockUser.ID,
		mockUser.Email,
		mockUser.Password,
	)

	query := "SELECT \\* FROM users WHERE email = \\$1 LIMIT 1;"
	mock.ExpectQuery(query).WillReturnRows(rows)
	user, err := userRepo.FindByEmail(mockUser.Email)
	if err != nil {
		t.Error(err)
	}

	if user.ID != mockUser.ID {
		t.Errorf("Error: wrong return, expected %d, found %d", mockUser.ID, user.ID)
	}
}
