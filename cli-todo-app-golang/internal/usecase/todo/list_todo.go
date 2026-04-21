package todo

import (
    "fmt"
    "context"
	"errors"

    "cli-todo-app-golang/internal/domain/repository"
    "cli-todo-app-golang/internal/interface/mapper"
	"cli-todo-app-golang/internal/domain/entity"
)

type ListTodoUseCase struct {
    todoRepo    repository.TodoRepository
    sessionRepo repository.SessionRepository
    userLogRepo repository.UserLogRepository
}

func NewListTodoUseCase(
    todoRepo repository.TodoRepository,
    sessionRepo repository.SessionRepository,
    userLogRepo repository.UserLogRepository,
) *ListTodoUseCase {
    return &ListTodoUseCase{
        todoRepo:    todoRepo,
        sessionRepo: sessionRepo,
        userLogRepo: userLogRepo,
    }
}

func (uc *ListTodoUseCase) Execute(ctx context.Context) ([]entity.Todo, error) {
	session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return nil, errors.New("you are not login!")
    }

	todos, err := uc.todoRepo.GetTodoByUserId(ctx, session.UserId)
	if err != nil {
		if errors.Is(err, repository.ErrTodoNotFound) {
			return nil, repository.ErrTodoNotFound
		}
		return nil, fmt.Errorf("failed to delete todo: %w", err)
	}

	action := "User Get Todo!"
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return nil, fmt.Errorf("failed to save log: %w", err)
	}

	return todos, nil
}
