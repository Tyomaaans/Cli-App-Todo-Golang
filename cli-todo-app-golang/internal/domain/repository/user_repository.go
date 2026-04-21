package repository

import (
    "errors"
    "context"
    
    "cli-todo-app-golang/internal/domain/entity"
	"cli-todo-app-golang/internal/interface/dto"
)

var (
	ErrUserNotFound      = errors.New("user not found!")
	ErrUserAlreadyExists = errors.New("user already exists! please login!")
)

type UserRepository interface {
    GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)

	ForgotPassword(ctx context.Context, email, password string) error
	UpdateName(ctx context.Context, username, name string) error
	UpdateEmail(ctx context.Context, username, email string) error
	UpdateUsername(ctx context.Context, username, newusername string) error
	UpdateUser(ctx context.Context, userId string, req dto.UpdateUserRequest) error
	DeleteUser(ctx context.Context, userId string) error

	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error

	GetAll(ctx context.Context) ([]entity.User, error)
}