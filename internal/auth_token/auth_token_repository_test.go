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

func Test_authTokenRepository_CreateAuthToken(t *testing.T) {
	type args struct {
		ctx       context.Context
		authToken domain.AuthToken
	}

	creationTime := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)
	expirationTime := creationTime.Add(24 * time.Hour)

	tests := []struct {
		name    string
		args    args
		mock    func(ts authTokenRepositoryTestSuite)
		want    int
		wantErr bool
	}{
		{
			name: "PASS - 토큰 생성",
			args: args{
				ctx: context.Background(),
				authToken: domain.AuthToken{
					UserID:         1,
					JwtToken:       "jwt_token",
					CreationTime:   creationTime,
					ExpirationTime: expirationTime,
				},
			},
			mock: func(ts authTokenRepositoryTestSuite) {
				ts.sqlMock.ExpectExec("INSERT INTO auth_tokens").
					WithArgs(1, "jwt_token", creationTime, expirationTime, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupAuthTokenRepositoryTestSuite()
			tt.mock(ts)

			// when
			got, err := ts.authTokenRepository.CreateAuthToken(tt.args.ctx, tt.args.authToken)

			// then
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_authTokenRepository_FindAuthTokenByUserIDAndJwtToken(t *testing.T) {
	type args struct {
		ctx    context.Context
		params domain.FindByUserIDAndJwtTokenParams
	}

	creationTime := time.Date(2023, time.June, 10, 0, 0, 0, 0, time.UTC)
	expirationTime := creationTime.Add(24 * time.Hour)

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
				query := "SELECT id, user_id, jwt_token, creation_time, expiration_time, active FROM auth_tokens"
				columns := []string{"id", "user_id", "jwt_token", "creation_time", "expiration_time", "active"}
				rows := sqlmock.NewRows(columns).AddRow(1, 1, "jwt_token", creationTime, expirationTime, true)
				ts.sqlMock.ExpectQuery(query).WithArgs(1, "jwt_token").WillReturnRows(rows)
			},
			want: domain.AuthToken{
				Base: domain.Base{
					ID: 1,
				},
				UserID:         1,
				JwtToken:       "jwt_token",
				CreationTime:   creationTime,
				ExpirationTime: expirationTime,
				Active:         true,
			},
			wantErr: false,
		},
		{
			name: "FAIL - 존재하지 않는 토큰 조회",
			args: args{
				ctx: context.Background(),
				params: domain.FindByUserIDAndJwtTokenParams{
					UserID:   1,
					JwtToken: "jwt_token",
				},
			},
			mock: func(ts authTokenRepositoryTestSuite) {
				query := "SELECT id, user_id, jwt_token, creation_time, expiration_time, active FROM `auth_tokens`"
				ts.sqlMock.ExpectQuery(query).WithArgs(1, "jwt_token").WillReturnError(sql.ErrNoRows)
			},
			want:    domain.AuthToken{},
			wantErr: true,
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

func Test_authTokenRepository_DeactivateAuthToken(t *testing.T) {
	type args struct {
		ctx    context.Context
		params domain.DeactivateAuthTokenParams
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts authTokenRepositoryTestSuite)
		want    int
		wantErr bool
	}{
		{
			name: "PASS - 활성 상태 변경",
			args: args{
				ctx: context.Background(),
				params: domain.DeactivateAuthTokenParams{
					UserID:   1,
					JwtToken: "target_token",
				},
			},
			mock: func(ts authTokenRepositoryTestSuite) {
				ts.sqlMock.ExpectExec("UPDATE auth_tokens SET active = 0").
					WithArgs(1, "target_token").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupAuthTokenRepositoryTestSuite()
			tt.mock(ts)

			// when
			err := ts.authTokenRepository.DeactivateAuthToken(tt.args.ctx, tt.args.params)

			// then
			if err := ts.sqlMock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
