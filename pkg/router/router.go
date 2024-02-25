package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"payhere/config"
	"payhere/docs"
)

func NewServeRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Payhere 백엔드 엔지니어 과제 REST API"
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
