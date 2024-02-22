package user

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"payhere/domain"
	"testing"
)

type userRepositoryTestSuite struct {
	sqlDB          *sql.DB
	sqlMock        sqlmock.Sqlmock
	userRepository domain.UserRepository
}

func setupUserRepositoryTestSuite() userRepositoryTestSuite {
	var us userRepositoryTestSuite

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	us.sqlDB = mockDB
	us.sqlMock = mock

	if err != nil {
		panic(err)
	}

	us.userRepository = NewUserRepository(mockDB)

	return us
}

func Test_userRepository_CreateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user domain.User
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts userRepositoryTestSuite)
		want    int
		wantErr bool
	}{
		{
			name: "PASS - 중복 되지 않은 유저 생성",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					MobileID: "01012345678",
					Password: "password",
					UseType:  domain.UserUseTypePlace,
				},
			},
			mock: func(ts userRepositoryTestSuite) {
				ts.sqlMock.ExpectExec("INSERT INTO `users`").
					WithArgs("01012345678", "password", "PLACE").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "PASS - 중복 된 유저 생성",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					MobileID: "01012345678",
					Password: "password",
					UseType:  domain.UserUseTypePlace,
				},
			},
			mock: func(ts userRepositoryTestSuite) {
				ts.sqlMock.ExpectExec("INSERT INTO `users`").
					WithArgs("01012345678", "password", "PLACE").
					WillReturnError(&mysql.MySQLError{Number: 1062, Message: "Duplicate entry"})
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserRepositoryTestSuite()
			tt.mock(ts)

			// when
			got, err := ts.userRepository.CreateUser(tt.args.ctx, tt.args.user)

			// then
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_userRepository_FindByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts userRepositoryTestSuite)
		want    *domain.User
		wantErr bool
	}{
		{
			name: "PASS - 존재하는 휴대폰번호로 조회",
			args: args{
				ctx:    context.Background(),
				userID: "01012345678",
			},
			mock: func(ts userRepositoryTestSuite) {
				query := "SELECT id, mobile_id, password, use_type FROM `users`"
				columns := []string{"id", "user_id", "password", "user_type"}
				rows := sqlmock.NewRows(columns).AddRow(1, "01012345678", "password", "PLACE")
				ts.sqlMock.ExpectQuery(query).WithArgs("01012345678").WillReturnRows(rows)
			},
			want: &domain.User{
				Base: domain.Base{
					ID: 1,
				},
				MobileID: "01012345678",
				Password: "password",
				UseType:  domain.UserUseTypePlace,
			},
			wantErr: false,
		},
		{
			name: "PASS - 존재하지 않는 휴대폰번호로 조회",
			args: args{
				ctx:    context.Background(),
				userID: "01012345678",
			},
			mock: func(ts userRepositoryTestSuite) {
				query := "SELECT id, mobile_id, password, use_type FROM `users`"
				ts.sqlMock.ExpectQuery(query).WithArgs("01012345678").WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserRepositoryTestSuite()
			tt.mock(ts)

			// when
			got, err := ts.userRepository.FindUserByMobileID(tt.args.ctx, tt.args.userID)

			// then
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
