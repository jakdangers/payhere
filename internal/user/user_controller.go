package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"payhere/domain"
	cerrors "payhere/pkg/cerror"
	"time"
)

func RegisterRoutes(e *gin.Engine, controller domain.UserController) {
	e.POST("/users", controller.CreateUser)
	e.POST("/users/login", controller.LoginUser)
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
