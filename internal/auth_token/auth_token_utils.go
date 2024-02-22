package auth_token

import (
	"github.com/golang-jwt/jwt/v5"
	"payhere/domain"
	"time"
)

func CreateAccessToken(user domain.User, secret string, exp time.Time) (accessToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    exp.Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
