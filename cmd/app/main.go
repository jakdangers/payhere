package main

import (
	"go.uber.org/fx"
	"payhere/config"
	"payhere/pkg/db"
	"payhere/pkg/httpserver"
	"payhere/pkg/router"
)

func main() {
	fx.New(
		// pkg module
		config.Module,
		db.SqlxModule,
		router.Module,

		// usecase module

		fx.Invoke(
			// handler Invoke
			// pkg Invoke
			httpserver.NewHTTPServer,
		),
	).Run()
}
