package user

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"payhere/domain"
	"payhere/mocks"
	"strings"
	"testing"
)

type userControllerTestSuite struct {
	router         *gin.Engine
	userService    *mocks.UserService
	userController domain.UserController
}

func setupUserControllerTestSuite(t *testing.T) userControllerTestSuite {
	var us userControllerTestSuite

	gin.SetMode(gin.TestMode)
	us.router = gin.Default()
	us.userService = mocks.NewUserService(t)
	us.userController = NewUserController(us.userService)
	RegisterRoutes(us.router, us.userController)

	return us
}

func Test_userController_CreateUser(t *testing.T) {
	tests := []struct {
		name  string
		input func() *bytes.Reader
		mock  func(ts userControllerTestSuite)
		code  int
	}{
		{
			name: "PASS - 휴대폰 번호, 하이픈 없음",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
				ts.userService.EXPECT().CreateUser(mock.Anything, domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "payhere",
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 휴대폰 번호, 하이픈 있음",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "010-1234-5678",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
				ts.userService.EXPECT().CreateUser(mock.Anything, domain.CreateUserRequest{
					UserID:   "010-1234-5678",
					Password: "payhere",
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "FAIL - 휴대폰 번호, 하이픈이 잘못됨",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "010-12345678",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 휴대폰 번호, 너무 짧음",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "0101234",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 휴대폰 번호, 잘못된 문자 포함",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "010-1234-abcd",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 휴대폰 번호 빈 문자열",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 휴대폰 번호, 잘못된 하이픈",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "0101234-5678",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {},
			code: http.StatusBadRequest,
		},
		{
			name: "PASS - 영어 소문자 한글자",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "p",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
				ts.userService.EXPECT().CreateUser(mock.Anything, domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "p",
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 영어 대문자 한글자",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "P",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
				ts.userService.EXPECT().CreateUser(mock.Anything, domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "P",
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 숫자 한글자",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "5",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
				ts.userService.EXPECT().CreateUser(mock.Anything, domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "5",
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 특수 기호 한글자",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "@",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
				ts.userService.EXPECT().CreateUser(mock.Anything, domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "@",
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "PASS - 255자 패스워드",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "payhere" + strings.Repeat("x", 248),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
				ts.userService.EXPECT().CreateUser(mock.Anything, domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "payhere" + strings.Repeat("x", 248),
				}).Return(nil).Once()
			},
			code: http.StatusNoContent,
		},
		{
			name: "FAIL – 0자 패스워드",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 256자 패스워드",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "payhere" + strings.Repeat("x", 249),
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {},
			code: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserControllerTestSuite(t)
			tt.mock(ts)
			req, _ := http.NewRequest(http.MethodPost, "/users", tt.input())
			req.Header.Set("Content-Type", "application/json")

			// when
			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			// then
			assert.Equal(t, tt.code, rec.Code)
			ts.userService.AssertExpectations(t)
		})
	}
}

func Test_userController_LoginUser(t *testing.T) {
	tests := []struct {
		name  string
		input func() *bytes.Reader
		mock  func(ts userControllerTestSuite)
		code  int
	}{
		{
			name: "PASS - 휴대폰 번호, 비밀번호",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
				ts.userService.EXPECT().LoginUser(mock.Anything, domain.LoginUserRequest{
					UserID:   "01012345678",
					Password: "payhere",
				}).
					Return(domain.LoginUserResponse{}, nil).Once()
			},
			code: http.StatusOK,
		},
		{
			name: "FAIL - 휴대폰 번호, 하이픈이 잘못됨",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "010-12345678",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 휴대폰 번호, 너무 짧음",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "0101234",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 휴대폰 번호, 잘못된 문자 포함",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "010-1234-abcd",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {
			},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 휴대폰 번호, 잘못된 하이픈",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "0101234-5678",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 휴대폰 번호 빈 문자열",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "",
					Password: "payhere",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {},
			code: http.StatusBadRequest,
		},
		{
			name: "FAIL - 비밀번호 빈 문자열",
			input: func() *bytes.Reader {
				req := domain.CreateUserRequest{
					UserID:   "01012345678",
					Password: "",
				}
				jsonData, _ := json.Marshal(req)

				return bytes.NewReader(jsonData)
			},
			mock: func(ts userControllerTestSuite) {},
			code: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			ts := setupUserControllerTestSuite(t)
			tt.mock(ts)
			req, _ := http.NewRequest(http.MethodPost, "/users/login", tt.input())
			req.Header.Set("Content-Type", "application/json")

			// when
			rec := httptest.NewRecorder()
			ts.router.ServeHTTP(rec, req)

			// then
			assert.Equal(t, tt.code, rec.Code)
			ts.userService.AssertExpectations(t)
		})
	}
}
