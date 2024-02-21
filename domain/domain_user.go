package domain

import (
	"context"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (int, error)
	FindByUserID(ctx context.Context, userID string) (*User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) error
}

type UserController interface {
	CreateUser(c *gin.Context)
}

type UserUseType string

const (
	UserUseTypePlace UserUseType = "PLACE"
)

type User struct {
	Base
	UserID   string
	Password string
	UseType  UserUseType
}
