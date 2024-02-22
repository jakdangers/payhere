package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"payhere/config"
	"payhere/domain"
	"payhere/internal/auth_token"
	cerrors "payhere/pkg/cerror"
	"regexp"
	"time"
)

var (
	mobileIDPattern     = regexp.MustCompile(`^(010-\d{4}-\d{4}|010\d{8})$`)
	hasHyphenPattern    = regexp.MustCompile(`-`)
	removeHyphenPattern = regexp.MustCompile(`-`)
)

type userService struct {
	userRepository domain.UserRepository
	authRepository domain.AuthTokenRepository
	cfg            *config.Config
}

func NewUserService(
	userRepository domain.UserRepository,
	authRepository domain.AuthTokenRepository,
	cfg *config.Config,
) *userService {
	return &userService{
		userRepository: userRepository,
		authRepository: authRepository,
		cfg:            cfg,
	}
}

var _ domain.UserService = (*userService)(nil)

func (us userService) CreateUser(ctx context.Context, req domain.CreateUserRequest) error {
	const op cerrors.Op = "user/service/createUser"

	phoneNumber, err := validateAndNormalizeMobileID(req.MobileID)
	if err != nil {
		return err
	}

	user, err := us.userRepository.FindUserByMobileID(ctx, phoneNumber)
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

	_, err = us.userRepository.CreateUser(ctx, domain.User{
		MobileID: phoneNumber,
		Password: hashedPassword,
		UseType:  domain.UserUseTypePlace,
	})
	if err != nil {
		return err
	}

	return nil
}

func (us userService) LoginUser(ctx context.Context, req domain.LoginUserRequest) (domain.LoginUserResponse, error) {
	const op cerrors.Op = "user/service/loginUser"

	mobileID, err := validateAndNormalizeMobileID(req.MobileID)
	if err != nil {
		return domain.LoginUserResponse{}, err
	}

	user, err := us.userRepository.FindUserByMobileID(ctx, mobileID)
	if err != nil {
		return domain.LoginUserResponse{}, err
	}
	if user == nil {
		return domain.LoginUserResponse{}, cerrors.E(op, cerrors.Invalid, "아이디 또는 비밀번호를 확인해주세요.")
	}

	if !compareHashAndPassword(req.Password, user.Password) {
		return domain.LoginUserResponse{}, cerrors.E(op, cerrors.Invalid, "아이디 또는 비밀번호를 확인해주세요.")
	}

	creationTime := time.Now().UTC()
	expirationTime := creationTime.Add(time.Hour * time.Duration(us.cfg.ExpiryHours))

	accessToken, err := auth_token.CreateAccessToken(*user, us.cfg.Auth.Secret, expirationTime)
	if err != nil {
		return domain.LoginUserResponse{}, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	_, err = us.authRepository.CreateAuthToken(ctx, domain.AuthToken{
		UserID:         user.ID,
		JwtToken:       accessToken,
		CreationTime:   creationTime,
		ExpirationTime: expirationTime,
	})
	if err != nil {
		return domain.LoginUserResponse{}, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return domain.LoginUserResponse{
		AccessToken: accessToken,
		ExpiresIn:   expirationTime.Unix(),
	}, nil
}

func (us userService) LogoutUser(ctx context.Context, req domain.LogoutUserRequest) error {
	const op cerrors.Op = "user/service/LogoutUser"

	authToken, err := us.authRepository.FindAuthTokenByUserIDAndJwtToken(ctx, domain.FindByUserIDAndJwtTokenParams{
		UserID:   req.UserID,
		JwtToken: req.AccessToken,
	})
	if !authToken.Active {
		return cerrors.E(op, cerrors.Invalid, "이미 로그아웃된 사용자입니다.")
	}
	if err != nil {
		return cerrors.E(op, err, "서버 에러가 발생했습니다.")
	}

	if err := us.authRepository.DeactivateAuthToken(ctx, domain.DeactivateAuthTokenParams{
		UserID:   req.UserID,
		JwtToken: req.AccessToken,
	}); err != nil {
		return cerrors.E(op, err, "서버 에러가 발생했습니다.")
	}

	return nil
}

func hashPasswordWithSalt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func compareHashAndPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func validateAndNormalizeMobileID(phoneNumber string) (string, error) {
	const op cerrors.Op = "user/service/validateAndNormalizeMobileID"

	if !mobileIDPattern.MatchString(phoneNumber) {
		return "", cerrors.E(op, cerrors.Invalid, "잘못된 휴대폰번호입니다.")
	}

	if hasHyphenPattern.MatchString(phoneNumber) {
		return removeHyphenPattern.ReplaceAllString(phoneNumber, ""), nil
	}

	return phoneNumber, nil
}
