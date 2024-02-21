package user

import (
	"context"
	"database/sql"
	"errors"
	"payhere/domain"
	cerrors "payhere/pkg/cerror"
)

type userRepository struct {
	sqlDB *sql.DB
}

func NewUserRepository(sqlDB *sql.DB) *userRepository {
	return &userRepository{
		sqlDB: sqlDB,
	}
}

var _ domain.UserRepository = (*userRepository)(nil)

func (u userRepository) CreateUser(ctx context.Context, user domain.User) (int, error) {
	const op cerrors.Op = "user/userRepository/createUser"

	result, err := u.sqlDB.ExecContext(ctx, createUserQuery, user.UserID, user.Password, user.UseType)
	if err != nil {
		return 0, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return int(userID), nil
}

func (u userRepository) FindByUserID(ctx context.Context, userID string) (*domain.User, error) {
	const op cerrors.Op = "user/userRepository/findByUserID"
	var user domain.User

	err := u.sqlDB.QueryRowContext(ctx, findByUserIDQuery, userID).
		Scan(&user.ID, &user.UserID, &user.Password, &user.UseType)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, cerrors.E(op, cerrors.Internal, err, "서버 에러가 발생했습니다.")
	}

	return &user, nil
}
