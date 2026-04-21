package auth

import (
	"fmt"
	"errors"
	"context"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"cli-todo-app-golang/internal/interface/dto"
	"cli-todo-app-golang/internal/interface/mapper"
    "cli-todo-app-golang/internal/domain/repository"
	appvalidator "cli-todo-app-golang/internal/validator"
)

type RegisterUseCase struct {
    userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	validate    *validator.Validate
}

func NewRegisterUseCase(
    userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	validate    *validator.Validate,
) *RegisterUseCase {
    return &RegisterUseCase{
        userRepo:    userRepo,
		sessionRepo: sessionRepo,
		validate: validate,
    }
}

func (uc *RegisterUseCase) Execute(ctx context.Context, req dto.RegisterRequest) error {
	if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	if req.Password != req.ConfirmPassword {
		return errors.New("password do not match")
	}

	if _, err := uc.sessionRepo.Get(ctx); err == nil {
		return errors.New("you are logged in! logout first to register!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := mapper.ToEntityFromRegisterRequest(string(hashedPassword), req)
	if err := uc.userRepo.Create(ctx, &user); err != nil {
		if errors.Is(err, repository.ErrUserAlreadyExists) {
			return repository.ErrUserAlreadyExists
		}
		return fmt.Errorf("create user failed: %w", err)
	}

	return nil
}