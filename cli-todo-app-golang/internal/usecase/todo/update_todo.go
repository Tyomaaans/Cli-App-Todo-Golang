package todo

import (
    "fmt"
    "context"
    "errors"
    "strings"

    "github.com/go-playground/validator/v10"

    "cli-todo-app-golang/internal/interface/mapper"
    "cli-todo-app-golang/internal/domain/repository"
    "cli-todo-app-golang/internal/interface/dto"
    appvalidator "cli-todo-app-golang/internal/validator"
)

type UpdateTodoUseCase struct {
    todoRepo    repository.TodoRepository
    sessionRepo repository.SessionRepository
    userLogRepo repository.UserLogRepository
    validate    *validator.Validate
}

func NewUpdateTodoUseCase(
    todoRepo repository.TodoRepository,
    sessionRepo repository.SessionRepository,
    userLogRepo repository.UserLogRepository,
    validate *validator.Validate,
) *UpdateTodoUseCase {
    return &UpdateTodoUseCase{
        todoRepo:    todoRepo,
        sessionRepo: sessionRepo,
        userLogRepo: userLogRepo,
        validate:    validate,
    }
}

func (uc *UpdateTodoUseCase) Execute(ctx context.Context, req dto.UpdateTodoRequest) error {
    if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
        return err
    }

    var update []string
    if req.Task == "" && req.Priority == "" && req.Progress == "" && req.DueDate == "" {
        return errors.New("update todo requires at least one field (Task, Priority, Progress, or Due Date)!")
    }
    if req.Task != "" {
        update = append(update, "Task")
    }
    if req.Priority != "" {
        update = append(update, "Priority")
    }
    if req.Progress != "" {
        update = append(update, "Progress")
    } 
    if req.DueDate != "" {
        update = append(update, "DueDate")
    }

    session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login! login first to update todo!")
    }

    if err := uc.todoRepo.UpdateTodo(ctx, session.UserId, req); err != nil {
		if errors.Is(err, repository.ErrTodoAlreadyExists) {
			return repository.ErrTodoAlreadyExists
		}
		return fmt.Errorf("failed to add todo: %w", err)
	}
    
    var results string
    if len(update) > 1 {
        results = strings.Join(update, ", ")
    }

    action := fmt.Sprintf("User Update %s Todo", results)
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

    return nil
}