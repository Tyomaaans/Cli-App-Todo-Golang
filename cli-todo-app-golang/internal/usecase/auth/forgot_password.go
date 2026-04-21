package auth

import (
	"errors"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"

	"cli-todo-app-golang/internal/domain/repository"
	"cli-todo-app-golang/internal/interface/dto"
	"cli-todo-app-golang/internal/interface/mapper"
	appvalidator "cli-todo-app-golang/internal/validator"
)

type ForgotPasswordUseCase struct {
	userRepo       repository.UserRepository
	sessionRepo    repository.SessionRepository
	userLogRepo    repository.UserLogRepository
	validate       *validator.Validate
}

func NewForgotPasswordUseCase (
	userRepo    repository.UserRepository,
	sessionRepo repository.SessionRepository,
	userLogRepo repository.UserLogRepository,
	validate    *validator.Validate,
) *ForgotPasswordUseCase {
	return &ForgotPasswordUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		userLogRepo: userLogRepo,
		validate:    validate,
	}
}

func (uc *ForgotPasswordUseCase) Execute(ctx context.Context, req dto.ForgotPasswordRequest) error {
	if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	if req.Password != req.ConfirmPassword {
		return errors.New("password do not match!")
	}

	if _, err := uc.sessionRepo.Get(ctx); err == nil {
		return errors.New("you are logged in! use update password!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed hash password!")
	}

	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return repository.ErrUserNotFound
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	if err := uc.userRepo.ForgotPassword(ctx, user.Email, string(hashedPassword)); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return repository.ErrUserNotFound
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	action  := "User Changed Password from Forgot Password!"
	userLog := mapper.ToEntityFromNewLogUser(user.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to log session: %w", err)
	}

	return nil
}