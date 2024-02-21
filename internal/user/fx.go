package user

import (
	"go.uber.org/fx"
	"payhere/domain"
)

var Module = fx.Module(
	"user", fx.Provide(
		fx.Annotate(NewUserRepository, fx.As(new(domain.UserRepository))),
		fx.Annotate(NewUserService, fx.As(new(domain.UserService))),
		fx.Annotate(NewUserController, fx.As(new(domain.UserController))),
	),
)
