package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"payhere/config"
	"payhere/internal/auth_token"
	"payhere/internal/product"
	"payhere/internal/user"
	"payhere/pkg/db"
	"payhere/pkg/router"
	"syscall"
	"time"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// infrastructure
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.NewSql(cfg)
	if err != nil {
		log.Fatal(err)
	}
	router := router.NewServeRouter(cfg)

	// domain
	authTokenRepository := auth_token.NewAuthTokenRepository(db)
	userRepsitory := user.NewUserRepository(db)
	productRepository := product.NewProductRepository(db)

	// service
	userService := user.NewUserService(userRepsitory, authTokenRepository, cfg)
	productService := product.NewProductService(userRepsitory, productRepository)

	// controller
	userController := user.NewUserController(userService)
	productController := product.NewProductController(productService)

	// routes
	user.RegisterRoutes(router, userController, authTokenRepository, cfg)
	product.RegisterRoutes(router, productController, authTokenRepository, cfg)

	// http server
	srv := &http.Server{Addr: cfg.HTTP.Port, Handler: router}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
