package todo

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

type DeleteTodoUseCase struct {
    todoRepo    repository.TodoRepository
    sessionRepo repository.SessionRepository
    userLogRepo repository.UserLogRepository
    validate    *validator.Validate
}

func NewDeleteTodoUseCase(
    todoRepo repository.TodoRepository,
    sessionRepo repository.SessionRepository,
    userLogRepo repository.UserLogRepository,
    validate *validator.Validate,
) *DeleteTodoUseCase {
    return &DeleteTodoUseCase{
        todoRepo:    todoRepo,
        sessionRepo: sessionRepo,
        userLogRepo: userLogRepo,
        validate:    validate,
    }
}

func (uc *DeleteTodoUseCase) Execute(ctx context.Context, req dto.DeleteTodoRequest) error {
	if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
        return err
    }

	session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login! login first to delete todo!")
    }

	if err := uc.todoRepo.DeleteTodo(ctx, req.DeleteIndex, session.UserId); err != nil {
		if errors.Is(err, repository.ErrTodoNotFound) {
			return repository.ErrTodoNotFound
		}
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	action := "User Deleted Todo!"
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

	return nil
}