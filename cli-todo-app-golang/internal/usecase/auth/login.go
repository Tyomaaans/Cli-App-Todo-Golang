package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"

	"cli-todo-app-golang/internal/domain/entity"
	"cli-todo-app-golang/internal/domain/repository"
	"cli-todo-app-golang/internal/interface/dto"
	"cli-todo-app-golang/internal/interface/mapper"
	appvalidator "cli-todo-app-golang/internal/validator"
)

type LoginUseCase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	userLogRepo repository.UserLogRepository
	validate    *validator.Validate
}

func NewLoginUseCase (
	userRepo    repository.UserRepository,
	sessionRepo repository.SessionRepository,
	userLogRepo repository.UserLogRepository,
	validate    *validator.Validate,
) *LoginUseCase {
	return &LoginUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		userLogRepo: userLogRepo,
		validate:    validate,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req dto.LoginRequest) error {
	var ErrInvalidCredentials = errors.New("username or password wrong")
	
	if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	if _, err := uc.sessionRepo.Get(ctx); err == nil {
		return errors.New("you are already login!")
	}

	var user  *entity.User
	var err    error
	var action string

	if req.Username != "" {
		user, err = uc.userRepo.GetByUsername(ctx, req.Username)
		action = "User Login With Username!"
	} else {
		user, err = uc.userRepo.GetByEmail(ctx, req.Email)
		action = "User Login With Email!"
	}

	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrInvalidCredentials
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return ErrInvalidCredentials
	}

	newSession := mapper.ToEntityFromSession(user.UserId, user.Name, user.Username, user.Email, time.Now())
	if err := uc.sessionRepo.Session(ctx, &newSession); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	userLog := mapper.ToEntityFromNewLogUser(user.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

	return nil
}