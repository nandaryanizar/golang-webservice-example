package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httpHandlers "github.com/nandaryanizar/golang-webservice-example/handlers/http"
	"github.com/sarulabs/di"

	"github.com/nandaryanizar/golang-webservice-example/entities"

	"github.com/nandaryanizar/golang-webservice-example/services/mocks"
)

func TestRequestToken(t *testing.T) {
	userService := new(mocks.UserServiceMock)
	userService.On("AuthenticateUser", "test.valid@gmail.com", "testpassword123").Return(entities.User{
		ID:    1,
		Email: "test.valid@gmail.com",
	}, nil)
	userService.On("AuthenticateUser", "test.invalid@gmail.com", "testpassword123").Return(entities.User{
		ID:    0,
		Email: "test.invalid@gmail.com",
	}, errors.New("Invalid username and/or password"))

	builder, _ := di.NewBuilder()
	builder.Add(di.Def{
		Name: "user-service",
		Build: func(ctn di.Container) (interface{}, error) {
			return userService, nil
		},
	})
	app := builder.Build()

	user := entities.User{
		Email:    "test.valid@gmail.com",
		Password: "testpassword123",
	}
	u, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", "/token", strings.NewReader(string(u)))
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	h := http.HandlerFunc(httpHandlers.Token)
	handlerWithDiMiddleware := di.HTTPMiddleware(h, app, func(msg string) {
		t.Error(msg)
	})
	handlerWithDiMiddleware.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler return wrong status code, expected %d, found %d", status, http.StatusOK)
	}
}

func TestFailedRequestToken(t *testing.T) {
	userService := new(mocks.UserServiceMock)
	userService.On("AuthenticateUser", "test.valid@gmail.com", "").Return(entities.User{
		ID:    0,
		Email: "test.valid@gmail.com",
	}, errors.New("Invalid username and/or password"))
	userService.On("AuthenticateUser", "", "testpassword123").Return(entities.User{
		ID:    0,
		Email: "test.valid@gmail.com",
	}, errors.New("Invalid username and/or password"))
	userService.On("AuthenticateUser", "test.invalid@gmail.com", "testpassword123").Return(entities.User{
		ID:    0,
		Email: "test.invalid@gmail.com",
	}, errors.New("Invalid username and/or password"))

	builder, _ := di.NewBuilder()
	builder.Add(di.Def{
		Name: "user-service",
		Build: func(ctn di.Container) (interface{}, error) {
			return userService, nil
		},
	})
	app := builder.Build()

	h := http.HandlerFunc(httpHandlers.Token)
	handlerWithDiMiddleware := di.HTTPMiddleware(h, app, func(msg string) {
		t.Error(msg)
	})

	cases := []struct {
		email    string
		password string
		status   int
	}{
		{"test.valid@gmail.com", "", http.StatusUnauthorized},
		{"", "testpassword123", http.StatusUnauthorized},
		{"test.invalid@gmail.com", "testpassword123", http.StatusUnauthorized},
	}

	for _, tc := range cases {
		user := entities.User{
			Email:    tc.email,
			Password: tc.password,
		}
		u, err := json.Marshal(user)
		if err != nil {
			t.Error(err)
		}

		req, err := http.NewRequest("POST", "/token", strings.NewReader(string(u)))
		if err != nil {
			t.Error(err)
		}

		rec := httptest.NewRecorder()
		handlerWithDiMiddleware.ServeHTTP(rec, req)

		if status := rec.Code; status != tc.status {
			t.Errorf("Handler return wrong status code, expected %d, found %d", status, tc.status)
		}
	}
}
