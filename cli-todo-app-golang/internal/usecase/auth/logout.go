package auth

import (
	"fmt"
	"context"
	"errors"

    "cli-todo-app-golang/internal/domain/repository"
	"cli-todo-app-golang/internal/interface/mapper"
)

type LogoutUseCase struct {
	sessionRepo repository.SessionRepository
	userLogRepo repository.UserLogRepository
}

func NewLogoutUseCase (
	sessionRepo repository.SessionRepository,
	userLogRepo repository.UserLogRepository,
) *LogoutUseCase {
	return &LogoutUseCase {
		sessionRepo: sessionRepo,
		userLogRepo: userLogRepo,
	}
}

func (uc *LogoutUseCase) Execute(ctx context.Context) error {
	session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login!")
    }

	if err := uc.sessionRepo.ClearSession(ctx); err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}

	action := "User Logout!"
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

	return nil
}
