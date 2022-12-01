package di

import (
	"github.com/akbaralishaikh/denti/pkg/config"
	"github.com/akbaralishaikh/denti/pkg/logger"
	"github.com/akbaralishaikh/denti/pkg/login"
	"github.com/akbaralishaikh/denti/pkg/storage"
	"github.com/akbaralishaikh/denti/pkg/storage/orm"
	"github.com/akbaralishaikh/denti/pkg/user"

	"go.uber.org/dig"
)

var container = dig.New()

func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)

	// DB
	container.Provide(storage.NewDb)

	// logger
	container.Provide(logger.NewLogger)

	// login
	container.Provide(orm.NewLoginRepo)
	container.Provide(login.NewLoginService)

	// user
	container.Provide(orm.NewUserRepo)
	container.Provide(user.NewUserService)
	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
