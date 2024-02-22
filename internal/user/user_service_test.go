package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"payhere/config"
	"payhere/domain"
	"payhere/mocks"
	"testing"
)

type userServiceTestSuite struct {
	userRepository      *mocks.UserRepository
	authTokenRepository *mocks.AuthTokenRepository
	service             domain.UserService
}

func setupUserServiceTestSuite(t *testing.T) userServiceTestSuite {
	var us userServiceTestSuite

	us.userRepository = mocks.NewUserRepository(t)
	us.authTokenRepository = mocks.NewAuthTokenRepository(t)
	us.service = NewUserService(us.userRepository, us.authTokenRepository, &config.Config{
		Auth: config.Auth{
			Secret:      "test_secret",
			ExpiryHours: 24,
		},
	})

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
					MobileID: "01012345678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.userRepository.EXPECT().FindUserByMobileID(mock.Anything, "01012345678").Return(nil, nil).Once()
				ts.userRepository.EXPECT().CreateUser(mock.Anything, mock.MatchedBy(func(user domain.User) bool {
					return user.MobileID == "01012345678" && compareHashAndPassword("payhere", user.Password)
				})).Return(1, nil).Once()
			},
			wantErr: false,
		},
		{
			name: "PASS - 중복되지 않는 하이픈 있는 휴대폰 번호로 사용자 생성",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					MobileID: "010-1234-5678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.userRepository.EXPECT().FindUserByMobileID(mock.Anything, "01012345678").Return(nil, nil).Once()
				ts.userRepository.EXPECT().CreateUser(mock.Anything, mock.MatchedBy(func(user domain.User) bool {
					return user.MobileID == "01012345678" && compareHashAndPassword("payhere", user.Password)
				})).Return(1, nil).Once()
			},
			wantErr: false,
		},
		{
			name: "FAIL - 중복되는 하이픈 없는 휴대폰 번호로 사용자 생성",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					MobileID: "01012345678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.userRepository.EXPECT().FindUserByMobileID(mock.Anything, "01012345678").Return(&domain.User{}, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "FAIL - 중복되는 하이픈 있는 휴대폰 번호로 사용자 생성",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					MobileID: "010-1234-5678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.userRepository.EXPECT().FindUserByMobileID(mock.Anything, "01012345678").Return(&domain.User{}, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "FAIL - 유효하지 않은 번호, 하이픈이 잘못됨",
			args: args{
				ctx: context.Background(),
				req: domain.CreateUserRequest{
					MobileID: "010-12345678",
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
					MobileID: "0101234",
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
					MobileID: "010-1234-abcd",
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
					MobileID: "",
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
					MobileID: "0101234-5678",
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

func Test_userService_LoginUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req domain.LoginUserRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts userServiceTestSuite)
		wantErr bool
	}{
		{
			name: "PASS - 유효한 휴대폰번호, 패스워드",
			args: args{
				ctx: context.Background(),
				req: domain.LoginUserRequest{
					MobileID: "01012345678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				hashPassword, _ := hashPasswordWithSalt("payhere")
				ts.userRepository.EXPECT().FindUserByMobileID(mock.Anything, "01012345678").
					Return(&domain.User{
						Base: domain.Base{
							ID: 1,
						},
						MobileID: "01012345678",
						Password: hashPassword,
						UseType:  domain.UserUseTypePlace,
					}, nil).Once()
				ts.authTokenRepository.EXPECT().CreateAuthToken(mock.Anything, mock.Anything).
					Return(1, nil).Once()
			},
			wantErr: false,
		},
		{
			name: "FAIL - 유효한 휴대폰 번호 잘못된 패스워드",
			args: args{
				ctx: context.Background(),
				req: domain.LoginUserRequest{
					MobileID: "01012345678",
					Password: "wrong_payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				hashPassword, _ := hashPasswordWithSalt("payhere")
				ts.userRepository.EXPECT().FindUserByMobileID(mock.Anything, "01012345678").
					Return(&domain.User{
						Base: domain.Base{
							ID: 1,
						},
						MobileID: "01012345678",
						Password: hashPassword,
						UseType:  domain.UserUseTypePlace,
					}, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "FAIL - 유효하지 않은 휴대폰 번호",
			args: args{
				ctx: context.Background(),
				req: domain.LoginUserRequest{
					MobileID: "01012345678",
					Password: "payhere",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.userRepository.EXPECT().FindUserByMobileID(mock.Anything, "01012345678").
					Return(nil, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "FAIL - 유효하지 않은 번호, 하이픈이 잘못됨",
			args: args{
				ctx: context.Background(),
				req: domain.LoginUserRequest{
					MobileID: "010-12345678",
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
				req: domain.LoginUserRequest{
					MobileID: "0101234",
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
				req: domain.LoginUserRequest{
					MobileID: "010-1234-abcd",
					Password: "payhere",
				},
			},
			mock:    func(ts userServiceTestSuite) {},
			wantErr: true,
		},
		{
			name: "FAIL - 빈 휴대폰 번호",
			args: args{
				ctx: context.Background(),
				req: domain.LoginUserRequest{
					MobileID: "",
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
				req: domain.LoginUserRequest{
					MobileID: "0101234-5678",
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
			_, err := ts.service.LoginUser(tt.args.ctx, tt.args.req)

			// then
			ts.userRepository.AssertExpectations(t)
			ts.authTokenRepository.AssertExpectations(t)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_userService_LogoutUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req domain.LogoutUserRequest
	}

	tests := []struct {
		name    string
		args    args
		mock    func(ts userServiceTestSuite)
		wantErr bool
	}{
		{
			name: "PASS - 로그아웃 성공",
			args: args{
				ctx: context.Background(),
				req: domain.LogoutUserRequest{
					UserID:      1,
					AccessToken: "access_token",
				},
			},
			mock: func(ts userServiceTestSuite) {
				ts.authTokenRepository.EXPECT().FindAuthTokenByUserIDAndJwtToken(mock.Anything, domain.FindByUserIDAndJwtTokenParams{
					UserID:   1,
					JwtToken: "access_token",
				}).Return(domain.AuthToken{
					UserID:   1,
					JwtToken: "access_token",
					Active:   true,
				}, nil).Once()
				ts.authTokenRepository.EXPECT().DeactivateAuthToken(mock.Anything, domain.DeactivateAuthTokenParams{
					UserID:   1,
					JwtToken: "access_token",
				}).Return(nil).Once()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserServiceTestSuite(t)
			tt.mock(ts)

			// when
			err := ts.service.LogoutUser(tt.args.ctx, tt.args.req)

			// then
			ts.userRepository.AssertExpectations(t)
			ts.authTokenRepository.AssertExpectations(t)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}

func Test_validateAndNormalizeMobileID(t *testing.T) {
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
			got, err := validateAndNormalizeMobileID(tt.input)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Equalf(t, tt.wantErr, err != nil, err.Error())
			}
		})
	}
}
