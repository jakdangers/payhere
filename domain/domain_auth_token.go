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
	LastAccessTime time.Time
	IsLoggedOut    bool
}

type AuthTokenRepository interface {
	FindAuthTokenByUserIDAndJwtToken(ctx context.Context, params FindByUserIDAndJwtTokenParams) (AuthToken, error)
}
