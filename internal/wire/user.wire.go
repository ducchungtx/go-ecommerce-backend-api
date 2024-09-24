//go:build wireinject

package wire

import (
	"goecommerce/internal/controller"
	"goecommerce/internal/repo"
	"goecommerce/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return &controller.UserController{}, nil
}
