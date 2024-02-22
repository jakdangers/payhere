package auth_token

import (
	"context"
	"database/sql"
	"payhere/domain"
	cerrors "payhere/pkg/cerror"
)

const (
	authTokenActive = 1
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

func (repo authTokenRepository) CreateAuthToken(ctx context.Context, token domain.AuthToken) (int, error) {
	const op cerrors.Op = "auth_token/authTokenRepository/CreateAuthToken"

	result, err := repo.sqlDB.ExecContext(
		ctx,
		createAuthTokenQuery,
		token.UserID,
		token.JwtToken,
		token.CreationTime,
		token.ExpirationTime,
		authTokenActive,
	)
	if err != nil {
		return 0, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	tokenID, err := result.LastInsertId()
	if err != nil {
		return 0, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return int(tokenID), nil
}

func (repo authTokenRepository) DeactivateAuthToken(ctx context.Context, params domain.DeactivateAuthTokenParams) error {
	const op cerrors.Op = "auth_token/authTokenRepository/DeactivateAuthToken"

	_, err := repo.sqlDB.ExecContext(ctx, deactivateAuthTokenQuery, params.UserID, params.JwtToken)
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return nil
}

func (repo authTokenRepository) FindAuthTokenByUserIDAndJwtToken(ctx context.Context, params domain.FindByUserIDAndJwtTokenParams) (domain.AuthToken, error) {
	const op cerrors.Op = "auth_token/authTokenRepository/FindByUserIDAndJwtToken"
	var token domain.AuthToken

	err := repo.sqlDB.QueryRowContext(ctx, findAuthTokenByUserIDAndJwtTokenQuery, params.UserID, params.JwtToken).
		Scan(&token.ID, &token.UserID, &token.JwtToken, &token.CreationTime, &token.ExpirationTime, &token.Active)
	if err != nil {
		return token, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return token, nil
}
