package registry

import (
	"github.com/nandaryanizar/fury"
	"github.com/nandaryanizar/golang-webservice-example/internal/app/logging"
	"github.com/nandaryanizar/golang-webservice-example/repositories"
	"github.com/nandaryanizar/golang-webservice-example/services"
	"github.com/sarulabs/di"
)

// Container struct
type Container struct {
	Ctn di.Container
}

// NewContainer factory method
func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add(diServices...); err != nil {
		return nil, err
	}

	return &Container{Ctn: builder.Build()}, nil
}

var diServices = []di.Def{
	{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return logging.Logger, nil
		},
	},
	{
		Name:  "fury-postgres",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return fury.Connect("./config/db/pg_database.yaml")
		},
		Close: func(obj interface{}) error {
			obj.(*fury.DB).Close()
			return nil
		},
	},
	{
		Name:  "user-repository",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return repositories.NewUserRepository(ctn.Get("fury-postgres").(*fury.DB)), nil
		},
	},
	{
		Name:  "user-service",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return services.NewUserService(ctn.Get("user-repository").(repositories.UserRepository)), nil
		},
	},
}
