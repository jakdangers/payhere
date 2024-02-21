package auth_token

import (
	"go.uber.org/fx"
	"payhere/domain"
)

var Module = fx.Module(
	"user", fx.Provide(
		fx.Annotate(NewAuthTokenRepository, fx.As(new(domain.AuthTokenRepository))),
	),
)
