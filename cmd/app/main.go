package main

import (
	"go.uber.org/fx"
	"payhere/config"
	"payhere/internal/auth_token"
	"payhere/internal/user"
	"payhere/pkg/db"
	"payhere/pkg/httpserver"
	"payhere/pkg/router"
)

func main() {
	fx.New(
		// pkg module
		config.Module,
		db.SqlModule,
		router.Module,

		// domain module
		user.Module,
		auth_token.Module,

		fx.Invoke(
			// routes Invoke
			user.RegisterRoutes,
			// pkg Invoke
			httpserver.NewHTTPServer,
		),
	).Run()
}
