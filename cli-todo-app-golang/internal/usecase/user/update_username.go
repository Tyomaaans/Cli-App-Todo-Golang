package user

import (
	"fmt"
	"context"
	"errors"

	"github.com/go-playground/validator/v10"

    "cli-todo-app-golang/internal/domain/repository"
	"cli-todo-app-golang/internal/interface/mapper"
	"cli-todo-app-golang/internal/interface/dto"
	appvalidator "cli-todo-app-golang/internal/validator"
)

type UpdateUsernameUseCase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	userLogRepo repository.UserLogRepository
	validate    *validator.Validate
}

func NewUpdateUsernameUseCase (
	userRepo    repository.UserRepository,
	sessionRepo repository.SessionRepository,
	userLogRepo repository.UserLogRepository,
	validate    *validator.Validate,
) *UpdateUsernameUseCase {
	return &UpdateUsernameUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		userLogRepo: userLogRepo,
		validate:    validate,
	}
}

func (uc *UpdateUsernameUseCase) Execute(ctx context.Context, req dto.UpdateUsernameRequest) error {
	if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login! login first to update username!")
    }

	if err := uc.userRepo.UpdateUsername(ctx, session.Username, req.Username); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return repository.ErrUserNotFound
		}
		return fmt.Errorf("failed to get user: %w", err)
	}
	
	action  := fmt.Sprintf("User Changed Username from %s to %s!", session.Username, req.Username)
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err  := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

	updateSession := mapper.ToEntityFromSession(session.UserId, session.Name, req.Username, session.Email, session.LoginTime)
	if err := uc.sessionRepo.Session(ctx, &updateSession); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}
