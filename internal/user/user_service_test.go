package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"payhere/domain"
	"payhere/mocks"
	"testing"
)

type userServiceTestSuite struct {
	repository *mocks.UserRepository
	service    domain.UserService
}

func setupUserServiceTestSuite(t *testing.T) userServiceTestSuite {
	var us userServiceTestSuite

	us.repository = mocks.NewUserRepository(t)
	us.service = NewUserService(us.repository)

	return us
}

func Test_userService_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req domain.CreateUserRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts userServiceTestSuite)
		wantErr bool
	}{
		{
			name: "PASS - 중복되지 않는 하이픈 없는 휴대폰 번호로 사용자 생성",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.repository.EXPECT().FindByUserID(mock.Anything, "01012345678").Return(nil, nil).Once()
				ts.repository.EXPECT().CreateUser(mock.Anything, mock.MatchedBy(func(user domain.User) bool {
					return user.UserID == "01012345678" && compareHashAndPassword("payhere", user.Password)
				})).Return(1, nil).Once()
			},
			wantErr: false,
		},
		{
			name: "PASS - 중복되지 않는 하이픈 있는 휴대폰 번호로 사용자 생성",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					UserID:   "010-1234-5678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.repository.EXPECT().FindByUserID(mock.Anything, "01012345678").Return(nil, nil).Once()
				ts.repository.EXPECT().CreateUser(mock.Anything, mock.MatchedBy(func(user domain.User) bool {
					return user.UserID == "01012345678" && compareHashAndPassword("payhere", user.Password)
				})).Return(1, nil).Once()
			},
			wantErr: false,
		},
		{
			name: "FAIL - 중복되는 하이픈 없는 휴대폰 번호로 사용자 생성",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.repository.EXPECT().FindByUserID(mock.Anything, "01012345678").Return(&domain.User{}, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "FAIL - 중복되는 하이픈 있는 휴대폰 번호로 사용자 생성",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					UserID:   "010-1234-5678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.repository.EXPECT().FindByUserID(mock.Anything, "01012345678").Return(&domain.User{}, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "FAIL - 유효하지 않은 번호, 하이픈이 잘못됨",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					UserID:   "010-12345678",
					Password: "payhere",
				},
			},
			mock:    func(ts userServiceTestSuite) {},
			wantErr: true,
		},
		{
			name: "FAIL - 유효하지 않은 번호, 너무 짧음",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					UserID:   "0101234",
					Password: "payhere",
				},
			},
			mock:    func(ts userServiceTestSuite) {},
			wantErr: true,
		},
		{
			name: "FAIL - 유효하지 않은 번호, 잘못된 문자 포함",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					UserID:   "010-1234-abcd",
					Password: "payhere",
				},
			},
			mock:    func(ts userServiceTestSuite) {},
			wantErr: true,
		},
		{
			name: "FAIL - 빈 문자열",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					UserID:   "",
					Password: "payhere",
				},
			},
			mock:    func(ts userServiceTestSuite) {},
			wantErr: true,
		},
		{
			name: "FAIL - 유효하지 않은 번호, 잘못된 하이픈",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					UserID:   "0101234-5678",
					Password: "payhere",
				},
			},
			mock:    func(ts userServiceTestSuite) {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserServiceTestSuite(t)
			tt.mock(ts)

			// when
			err := ts.service.CreateUser(tt.args.ctx, tt.args.req)

			// then
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_validateAndNormalizePhoneNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "PASS - 유효한 번호, 하이픈 없음", input: "01012345678", want: "01012345678", wantErr: false},
		{name: "PASS - 유효한 번호, 하이픈 있음", input: "010-1234-5678", want: "01012345678", wantErr: false},
		{name: "FAIL - 유효하지 않은 번호, 하이픈이 잘못됨", input: "010-12345678", want: "", wantErr: true},
		{name: "FAIL - 유효하지 않은 번호, 너무 짧음", input: "0101234", want: "", wantErr: true},
		{name: "FAIL - 유효하지 않은 번호, 잘못된 문자 포함", input: "010-1234-abcd", want: "", wantErr: true},
		{name: "FAIL - 빈 문자열", input: "", want: "", wantErr: true},
		{name: "FAIL - 유효하지 않은 번호, 잘못된 하이픈", input: "0101234-5678", want: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateAndNormalizePhoneNumber(tt.input)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
