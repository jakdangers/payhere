package auth_token

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"payhere/domain"
	"testing"
	"time"
)

type authTokenRepositoryTestSuite struct {
	sqlDB               *sql.DB
	sqlMock             sqlmock.Sqlmock
	authTokenRepository domain.AuthTokenRepository
}

func setupAuthTokenRepositoryTestSuite() authTokenRepositoryTestSuite {
	var ts authTokenRepositoryTestSuite

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	ts.sqlDB = mockDB
	ts.sqlMock = mock

	if err != nil {
		panic(err)
	}

	ts.authTokenRepository = NewAuthTokenRepository(mockDB)

	return ts
}

func Test_authTokenRepository_FindByUserIDAndJwtToken(t *testing.T) {
	type args struct {
		ctx    context.Context
		params domain.FindByUserIDAndJwtTokenParams
	}

	creationTime := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)
	expirationTime := creationTime.Add(24 * time.Hour)
	lastLoginTime := creationTime.Add(12 * time.Hour)
	//formattedCreationTime := creationTime.Format("2006-01-02 15:04:05")
	//formattedExpirationTime := expirationTime.Format("2006-01-02 15:04:05")
	//formattedLastLoginTime := lastLoginTime.Format("2006-01-02 15:04:05")

	tests := []struct {
		name    string
		args    args
		mock    func(ts authTokenRepositoryTestSuite)
		want    domain.AuthToken
		wantErr bool
	}{
		{
			name: "PASS - 토큰 조회",
			args: args{
				ctx: context.Background(),
				params: domain.FindByUserIDAndJwtTokenParams{
					UserID:   1,
					JwtToken: "jwt_token",
				},
			},
			mock: func(ts authTokenRepositoryTestSuite) {
				query := "SELECT id, user_id, jwt_token, creation_time, expiration_time, last_access_time, is_logged_out  FROM `auth_tokens` (.+)"
				columns := []string{"id", "user_id", "jwt_token", "creation_time", "expiration_time", "last_access_time", "is_logged_out"}
				rows := sqlmock.NewRows(columns).AddRow(1, 1, "jwt_token", creationTime, expirationTime, lastLoginTime, 0)
				ts.sqlMock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: domain.AuthToken{
				Base: domain.Base{
					ID: 1,
				},
				UserID:         1,
				JwtToken:       "jwt_token",
				CreationTime:   creationTime,
				ExpirationTime: expirationTime,
				LastAccessTime: lastLoginTime,
				IsLoggedOut:    false,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupAuthTokenRepositoryTestSuite()
			tt.mock(ts)

			// when
			got, err := ts.authTokenRepository.FindAuthTokenByUserIDAndJwtToken(tt.args.ctx, tt.args.params)

			// then
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
