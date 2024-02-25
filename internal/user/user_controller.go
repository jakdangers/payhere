package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"payhere/config"
	"payhere/domain"
	cerrors "payhere/pkg/cerrors"
	"payhere/pkg/router"
	"time"
)

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

// CreateUser
// @Tags User
// @Summary 회원가입
// @Description 사장님은 휴대폰 번호는 010-1234-5678, 01012345678 두개의 형식만 유효하고 비밀번호는 영문 대소문자, 숫자, 특수문자를 포함한 1자 이상 255자 이하의 문자열로 제한합니다.
// @Accept json
// @Produce json
// @Param CreateUserRequest body domain.CreateUserRequest true "회원가입 요청"
// @Success 204
// @Router /users [post]
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

// LoginUser
// @Tags User
// @Summary 로그인
// @Description 휴대폰 번호의 형식은 010-1234-5678, 01012345678 두개의 형식만 유효합니다
// @Accept json
// @Produce json
// @Param LoginUserRequest body domain.LoginUserRequest true "로그인 요청"
// @Success 200 {object} domain.LoginUserResponse
// @Router /users/login [post]
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

	c.JSON(domain.PayhereResponseFrom(http.StatusOK, res))
}

// LogoutUser
// @Tags User
// @Summary 로그아웃
// @Description 엑세스 토큰을 비활성화하고 로그아웃 처리합니다. (로그인 상태에서만 가능)
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 204
// @Router /users/logout [post]
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
