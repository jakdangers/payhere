package router

import (
	"github.com/gin-gonic/gin"
	cerrors "payhere/pkg/cerrors"
)

func GetUserIDFromContext(c *gin.Context) (int, error) {
	const op cerrors.Op = "router/GetUserIDFromContext"

	userID, ok := c.Get("userID")
	if !ok {
		return 0, cerrors.E(op, cerrors.Internal, "서버에 문제가 발생했습니다.")
	}

	userIDInt, ok := userID.(int)
	if !ok {
		return 0, cerrors.E(op, cerrors.Internal, "서버에 문제가 발생했습니다.")
	}

	return userIDInt, nil
}

func GetJWTTokenStringFromContext(c *gin.Context) (string, error) {
	const op cerrors.Op = "router/GetJWTTokenStringFromContext"

	tokenString, ok := c.Get("tokenString")
	if !ok {
		return "", cerrors.E(op, cerrors.Internal, "서버에 문제가 발생했습니다.")
	}

	userIDInt, ok := tokenString.(string)
	if !ok {
		return "", cerrors.E(op, cerrors.Internal, "서버에 문제가 발생했습니다.")
	}

	return userIDInt, nil
}
