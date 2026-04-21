package todo

import (
    "fmt"
    "context"
    "errors"

    "github.com/go-playground/validator/v10"

    "cli-todo-app-golang/internal/interface/mapper"
    "cli-todo-app-golang/internal/domain/repository"
    "cli-todo-app-golang/internal/interface/dto"
    appvalidator "cli-todo-app-golang/internal/validator"
)

type AddTodoUseCase struct {
    todoRepo    repository.TodoRepository
    sessionRepo repository.SessionRepository
    userLogRepo repository.UserLogRepository
    validate    *validator.Validate
}

func NewAddTodoUseCase(
    todoRepo repository.TodoRepository,
    sessionRepo repository.SessionRepository,
    userLogRepo repository.UserLogRepository,
    validate *validator.Validate,
) *AddTodoUseCase {
    return &AddTodoUseCase{
        todoRepo:    todoRepo,
        sessionRepo: sessionRepo,
        userLogRepo: userLogRepo,
        validate:    validate,
    }
}

func (uc *AddTodoUseCase) Execute(ctx context.Context, req dto.AddTodoRequest) error {
    if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
        return err
    }

    session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login! login first to add todo!")
    }

    newTodo := mapper.ToEntityFromAddTodoRequest(session.UserId, req)
    if err := uc.todoRepo.AddTodo(ctx, &newTodo); err != nil {
		if errors.Is(err, repository.ErrTodoAlreadyExists) {
			return repository.ErrTodoAlreadyExists
		}
		return fmt.Errorf("failed to add todo: %w", err)
	}

    action := "User Added Todo!"
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

    return nil
}