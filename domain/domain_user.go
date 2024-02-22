package domain

import (
	"context"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (int, error)
	FindUserByMobileID(ctx context.Context, userID string) (*User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) error
	LoginUser(ctx context.Context, req LoginUserRequest) (LoginUserResponse, error)
	LogoutUser(ctx context.Context, req LogoutUserRequest) error
}

type UserController interface {
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	LogoutUser(c *gin.Context)
}

type UserUseType string

const (
	UserUseTypePlace UserUseType = "PLACE"
)

type User struct {
	Base
	MobileID string
	Password string
	UseType  UserUseType
}
