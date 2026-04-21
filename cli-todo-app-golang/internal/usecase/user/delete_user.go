package user

import (
	"fmt"
	"context"
	"errors"

    "cli-todo-app-golang/internal/domain/repository"
	"cli-todo-app-golang/internal/interface/mapper"
)

type DeleteUserUseCase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	userLogRepo repository.UserLogRepository
}

func NewDeleteUserUseCase (
	userRepo    repository.UserRepository,
	sessionRepo repository.SessionRepository,
	userLogRepo repository.UserLogRepository,
) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		userLogRepo: userLogRepo,
	}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context) error {
	session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login! login first to delete todo!")
    }

	if err := uc.userRepo.DeleteUser(ctx, session.UserId); err != nil {
		if errors.Is(err, repository.ErrTodoNotFound) {
			return repository.ErrTodoNotFound
		}
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if err := uc.sessionRepo.ClearSession(ctx); err != nil {
		return fmt.Errorf("failed to clear session: %w", err)
	}

	action := "User Delete Account!"
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

	return nil
}