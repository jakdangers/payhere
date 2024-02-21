package auth_token

import (
	"context"
	"database/sql"
	"payhere/domain"
	cerrors "payhere/pkg/cerror"
)

type authTokenRepository struct {
	sqlDB *sql.DB
}

func NewAuthTokenRepository(sqlDB *sql.DB) *authTokenRepository {
	return &authTokenRepository{
		sqlDB: sqlDB,
	}
}

var _ domain.AuthTokenRepository = (*authTokenRepository)(nil)

func (repo authTokenRepository) FindAuthTokenByUserIDAndJwtToken(ctx context.Context, params domain.FindByUserIDAndJwtTokenParams) (domain.AuthToken, error) {
	const op cerrors.Op = "auth_token/authTokenRepository/findByUserIDAndJwtToken"
	var token domain.AuthToken

	err := repo.sqlDB.QueryRowContext(ctx, findAuthTokenByUserIDAndJwtTokenQuery, params.UserID, params.JwtToken).
		Scan(&token.ID, &token.UserID, &token.JwtToken, &token.CreationTime, &token.ExpirationTime, &token.LastAccessTime, &token.IsLoggedOut)
	if err != nil {
		return token, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return token, nil
}
