package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"payhere/domain"
	cerrors "payhere/pkg/cerrors"
	"strings"
	"time"
)

func parseJWTToken(tokenString string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func extractIDFromToken(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("Invalid token claims")
	}

	userIDClaim, ok := claims["userID"]
	if !ok {
		return 0, fmt.Errorf("UserID claim not found")
	}

	userID, ok := userIDClaim.(float64) // 여기서 적절한 타입으로 형변환 필요
	if !ok {
		return 0, fmt.Errorf("Invalid UserID type")
	}

	return int(userID), nil
}

func JWTMiddleware(secret string, authTokenRepository domain.AuthTokenRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(cerrors.NewSentinelAPIError(http.StatusUnauthorized, "로그인 후 이용해주세요"))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(cerrors.NewSentinelAPIError(http.StatusUnauthorized, "로그인 후 이용해주세요"))
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := parseJWTToken(tokenString, secret)
		if err != nil {
			c.JSON(cerrors.NewSentinelAPIError(http.StatusUnauthorized, "올바르지 않은 토큰입니다"))
			c.Abort()
			return
		}

		userID, err := extractIDFromToken(token)
		if err != nil {
			c.JSON(cerrors.NewSentinelAPIError(http.StatusUnauthorized, "올바르지 않은 토큰입니다"))
			c.Abort()
			return
		}

		authToken, err := authTokenRepository.FindAuthTokenByUserIDAndJwtToken(c, domain.FindByUserIDAndJwtTokenParams{
			UserID:   userID,
			JwtToken: tokenString,
		})
		if err != nil {
			c.JSON(cerrors.NewSentinelAPIError(http.StatusUnauthorized, "서버 에러가 발생했습니다."))
			c.Abort()
			return
		}
		if !authToken.Active {
			c.JSON(cerrors.NewSentinelAPIError(http.StatusUnauthorized, "로그인이 만료 되었습니다."))
			c.Abort()
		}
		if authToken.ExpirationTime.Before(time.Now().UTC()) {
			c.JSON(cerrors.NewSentinelAPIError(http.StatusUnauthorized, "로그인이 만료 되었습니다."))
			c.Abort()
		}

		c.Set("userID", userID)
		c.Set("tokenString", tokenString)

		c.Next()
	}
}
