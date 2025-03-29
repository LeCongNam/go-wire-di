// go:build wireinject
//go:build wireinject
// +build wireinject

package di_wire

import (
	"test/controllers"
	"test/repo"
	"test/services"

	"github.com/google/wire"
)

func InitUserRouterHandler() *controllers.UserController {
	panic(wire.Build(
		repo.NewUserRepo,
		services.NewUserService,
		controllers.NewUserController,
	))
}
