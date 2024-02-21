package main

import (
	"go.uber.org/fx"
	"payhere/config"
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

		fx.Invoke(
			// routes Invoke
			user.RegisterRoutes,
			// pkg Invoke
			httpserver.NewHTTPServer,
		),
	).Run()
}
