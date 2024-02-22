package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"payhere/config"
	"payhere/domain"
	cerrors "payhere/pkg/cerror"
	"payhere/pkg/router"
	"time"
)

// RegisterRoutes TODO : authRepository을 직접 사용하지 않도록 변경
func RegisterRoutes(e *gin.Engine, controller domain.UserController, authTokenRepository domain.AuthTokenRepository, cfg *config.Config) {
	api := e.Group("/users")
	{
		api.POST("", controller.CreateUser)
		api.POST("/login", controller.LoginUser)
		api.POST("/logout", router.JWTMiddleware(cfg.Auth.Secret, authTokenRepository), controller.LogoutUser)
	}
}

type userController struct {
	service domain.UserService
}

func NewUserController(service domain.UserService) *userController {
	return &userController{
		service: service,
	}
}

var _ domain.UserController = (*userController)(nil)

func (u userController) CreateUser(c *gin.Context) {
	var req domain.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	if err := u.service.CreateUser(ctx, req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.Status(http.StatusNoContent)
}

func (u userController) LoginUser(c *gin.Context) {
	var req domain.LoginUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	res, err := u.service.LoginUser(ctx, req)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (u userController) LogoutUser(c *gin.Context) {
	userID, err := router.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	tokenString, err := router.GetJWTTokenStringFromContext(c)
	if err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	if err := u.service.LogoutUser(ctx, domain.LogoutUserRequest{
		UserID:      userID,
		AccessToken: tokenString,
	}); err != nil {
		c.JSON(cerrors.ToSentinelAPIError(err))
		return
	}

	c.Status(http.StatusNoContent)
}
