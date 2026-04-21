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

type UpdateEmailUseCase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	userLogRepo repository.UserLogRepository
	validate    *validator.Validate
}

func NewUpdateEmailUseCase (
	userRepo    repository.UserRepository,
	sessionRepo repository.SessionRepository,
	userLogRepo repository.UserLogRepository,
	validate    *validator.Validate,
) *UpdateEmailUseCase {
	return &UpdateEmailUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		userLogRepo: userLogRepo,
		validate:    validate,
	}
}

func (uc *UpdateEmailUseCase) Execute(ctx context.Context, req dto.UpdateEmailRequest) error {
	if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
		return err
	}
	
	session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login! login first to update email!")
    }

	if err := uc.userRepo.UpdateEmail(ctx, session.Username, req.Email); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return repository.ErrUserNotFound
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	action := fmt.Sprintf("User Changed Email from %s to %s!", session.Email, req.Email)
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

	updateSession := mapper.ToEntityFromSession(session.UserId, session.Name, session.Username, req.Email, session.LoginTime)
	if err := uc.sessionRepo.Session(ctx, &updateSession); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}
