package domain

import (
	cerrors "payhere/pkg/cerror"
	"regexp"
)

var (
	mobileIDPattern = regexp.MustCompile(`^(010-\d{4}-\d{4}|010\d{8})$`)
	passwordPattern = regexp.MustCompile(`^[A-Za-z0-9@$!%*?&]{1,255}$`)
)

type CreateUserRequest struct {
	MobileID string `json:"mobileID"`
	Password string `json:"password"`
}

func (ur CreateUserRequest) Validate() error {
	const op cerrors.Op = "user/controller/valid"

	if !isValidMobileID(ur.MobileID) {
		return cerrors.E(op, cerrors.Invalid, "잘못된 휴대폰번호입니다.")
	}

	if !isValidPassword(ur.Password) {
		return cerrors.E(op, cerrors.Invalid, "잘못된 비밀번호입니다.")
	}

	return nil
}

// IsValidPhoneNumber
// 요구사항에 별도의 휴대폰번호 형식이 특정되어 있지 않아 하이픈이 있는 경우와 없는 경우를 모두 허용하도록 구현
func isValidMobileID(userID string) bool {
	return mobileIDPattern.MatchString(userID)
}

// IsValidPassword
// 패스워드는 영문 대소문자, 숫자, 특수문자를 포함한 1자 이상 255자 이하의 문자열로 제한
func isValidPassword(password string) bool {
	return passwordPattern.MatchString(password)
}

type LoginUserRequest struct {
	MobileID string `json:"mobileID"`
	Password string `json:"password"`
}

func (ur LoginUserRequest) Validate() error {
	const op cerrors.Op = "user/controller/valid"

	if !isValidMobileID(ur.MobileID) {
		return cerrors.E(op, cerrors.Invalid, "아이디 또는 비밀번호를 확인해주세요.")
	}

	if !isValidPassword(ur.Password) {
		return cerrors.E(op, cerrors.Invalid, "아이디 또는 비밀번호를 확인해주세요.")
	}

	return nil
}

type LoginUserResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
}

type LogoutUserRequest struct {
	UserID      int    `json:"userID"`
	AccessToken string `json:"accessToken"`
}
