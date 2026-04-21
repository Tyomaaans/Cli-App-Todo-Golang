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

type UpdateNameUseCase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	userLogRepo repository.UserLogRepository
	validate    *validator.Validate
}

func NewUpdateNameUseCase (
	userRepo    repository.UserRepository,
	sessionRepo repository.SessionRepository,
	userLogRepo repository.UserLogRepository,
	validate    *validator.Validate,
) *UpdateNameUseCase {
	return &UpdateNameUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		userLogRepo: userLogRepo,
		validate:    validate,
	}
}

func (uc *UpdateNameUseCase) Execute(ctx context.Context, req dto.UpdateNameRequest) error {
	if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}

	session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login! login first to update name!")
    }

	if err := uc.userRepo.UpdateName(ctx, session.Username, req.Name); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return repository.ErrUserNotFound
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	action  := fmt.Sprintf("User Changed Name from %s to %s!", session.Name, req.Name)
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err  := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

	updateSession := mapper.ToEntityFromSession(session.UserId, req.Name, session.Username, session.Email, session.LoginTime)
	if err := uc.sessionRepo.Session(ctx, &updateSession); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}
