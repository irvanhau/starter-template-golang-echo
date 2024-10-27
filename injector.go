//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"starter-template/configs"
	"starter-template/features/users"
	userHandler "starter-template/features/users/handler"
	userRepo "starter-template/features/users/repository"
	userService "starter-template/features/users/service"
	"starter-template/helper/jwt"
	"starter-template/routes"
	"starter-template/server"
	"starter-template/utils/database"
)

var userSet = wire.NewSet(
	userRepo.NewRepository,
	wire.Bind(new(users.UserDataInterface), new(*userRepo.UserData)),

	userService.NewService,
	wire.Bind(new(users.UserServiceInterface), new(*userService.UserService)),

	userHandler.NewHandler,
	wire.Bind(new(users.UserHandlerInterface), new(*userHandler.UserHandler)),
)

func InitializedServer() *server.Server {
	wire.Build(
		configs.InitConfig,
		database.InitDB,
		jwt.NewJWT,
		// DONT CHANGE THE FILE ABOVE

		userSet,

		// DONT CHANGE THE FILE BELOW
		routes.NewRoute,
		server.InitServer,
	)

	return nil
}
