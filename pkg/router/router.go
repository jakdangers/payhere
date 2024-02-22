package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"net/http"
	"payhere/config"
)

var Module = fx.Module("router", fx.Provide(NewServeRouter))

func NewServeRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	//r.Use(JWTMiddleware(cfg.Secret))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
