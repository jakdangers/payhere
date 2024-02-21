package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"payhere/domain"
	cerrors "payhere/pkg/cerror"
	"regexp"
)

var (
	phoneNumberPattern  = regexp.MustCompile(`^(010-\d{4}-\d{4}|010\d{8})$`)
	hasHyphenPattern    = regexp.MustCompile(`-`)
	removeHyphenPattern = regexp.MustCompile(`-`)
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

var _ domain.UserService = (*userService)(nil)

func (u userService) CreateUser(ctx context.Context, req domain.CreateUserRequest) error {
	const op cerrors.Op = "user/service/createUser"

	phoneNumber, err := validateAndNormalizePhoneNumber(req.UserID)
	if err != nil {
		return err
	}

	user, err := u.userRepository.FindByUserID(ctx, phoneNumber)
	if err != nil {
		return err
	}
	if user != nil {
		return cerrors.E(op, cerrors.Invalid, "이미 사용중인 휴대폰번호입니다.")
	}

	hashedPassword, err := hashPasswordWithSalt(req.Password)
	if err != nil {
		return cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	_, err = u.userRepository.CreateUser(ctx, domain.User{
		UserID:   phoneNumber,
		Password: hashedPassword,
		UseType:  domain.UserUseTypePlace,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u userService) LoginUser(ctx context.Context, req domain.LoginUserRequest) (domain.LoginUserResponse, error) {
	const op cerrors.Op = "user/service/createUser"

	phoneNumber, err := validateAndNormalizePhoneNumber(req.UserID)
	if err != nil {
		return domain.LoginUserResponse{}, err
	}

	user, err := u.userRepository.FindByUserID(ctx, phoneNumber)
	if err != nil {
		return domain.LoginUserResponse{}, err
	}
	if user == nil {
		return domain.LoginUserResponse{}, cerrors.E(op, cerrors.Invalid, "아이디 또는 비밀번호를 확인해주세요.")
	}

	if !compareHashAndPassword(req.Password, user.Password) {
		return domain.LoginUserResponse{}, cerrors.E(op, cerrors.Invalid, "아이디 또는 비밀번호를 확인해주세요.")
	}
}

func hashPasswordWithSalt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func compareHashAndPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func validateAndNormalizePhoneNumber(phoneNumber string) (string, error) {
	const op cerrors.Op = "user/service/validateAndNormalizePhoneNumber"

	if !phoneNumberPattern.MatchString(phoneNumber) {
		return "", cerrors.E(op, cerrors.Invalid, "잘못된 휴대폰번호입니다.")
	}

	if hasHyphenPattern.MatchString(phoneNumber) {
		return removeHyphenPattern.ReplaceAllString(phoneNumber, ""), nil
	}

	return phoneNumber, nil
}
