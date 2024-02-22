package domain

import (
	"context"
	"time"
)

type AuthToken struct {
	Base
	UserID         int
	JwtToken       string
	CreationTime   time.Time
	ExpirationTime time.Time
	Active         bool
}

type AuthTokenRepository interface {
	CreateAuthToken(ctx context.Context, token AuthToken) (int, error)
	DeactivateAuthToken(ctx context.Context, params DeactivateAuthTokenParams) error
	FindAuthTokenByUserIDAndJwtToken(ctx context.Context, params FindByUserIDAndJwtTokenParams) (AuthToken, error)
}
