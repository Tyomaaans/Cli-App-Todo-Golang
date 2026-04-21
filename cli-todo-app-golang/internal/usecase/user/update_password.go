package user

import (
	"fmt"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"

    "cli-todo-app-golang/internal/domain/repository"
	"cli-todo-app-golang/internal/interface/mapper"
	"cli-todo-app-golang/internal/interface/dto"
	appvalidator "cli-todo-app-golang/internal/validator"
)

type UpdatePasswordUseCase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	userLogRepo repository.UserLogRepository
	validate    *validator.Validate
}

func NewUpdatePasswordUseCase (
	userRepo    repository.UserRepository,
	sessionRepo repository.SessionRepository,
	userLogRepo repository.UserLogRepository,
	validate    *validator.Validate,
) *UpdatePasswordUseCase {
	return &UpdatePasswordUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		userLogRepo: userLogRepo,
		validate:    validate,
	}
}

func (uc *UpdatePasswordUseCase) Execute(ctx context.Context, req dto.UpdatePasswordRequest) error {
	if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	if req.Password != req.ConfirmPassword {
		return errors.New("password do not match!")
	}

	session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login! use forgot password!")
    }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed hash password!")
	}

	if err := uc.userRepo.ForgotPassword(ctx, session.Email, string(hashedPassword)); err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	action  := "User Changed Password from Update Password!"
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err  := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

	updateSession := mapper.ToEntityFromSession(session.UserId, session.Name, session.Username, session.Email, session.LoginTime)
	if err := uc.sessionRepo.Session(ctx, &updateSession); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}
