package user

import (
	"fmt"
	"context"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"

    "cli-todo-app-golang/internal/domain/repository"
	"cli-todo-app-golang/internal/interface/mapper"
	"cli-todo-app-golang/internal/interface/dto"
	appvalidator "cli-todo-app-golang/internal/validator"
)

type UpdateUserUseCase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	userLogRepo repository.UserLogRepository
	validate    *validator.Validate
}

func NewUpdateUserUseCase (
	userRepo    repository.UserRepository,
	sessionRepo repository.SessionRepository,
	userLogRepo repository.UserLogRepository,
	validate    *validator.Validate,
) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		userLogRepo: userLogRepo,
		validate:    validate,
	}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, req dto.UpdateUserRequest) error {
    if err := appvalidator.ValidateStruct(uc.validate, req); err != nil {
        return err
    }

	session, err := uc.sessionRepo.Get(ctx)
    if err != nil {
        return errors.New("you are not login! login first to update user!")
    }

    var update []string
    if req.Name == "" && req.Username == "" && req.Email == "" && req.Password == "" {
        return errors.New("update user requires at least one field (Name, USername, Email, or Password)!")
    }
    if req.Name != "" {
        update = append(update, "Name")
		session.Name = req.Name
    }
    if req.Username != "" {
        update = append(update, "Username")
		session.Username = req.Username
    }
    if req.Email != "" {
        update = append(update, "Email")
		session.Email = req.Email
    } 
    if req.Password != "" {
        update = append(update, "Password")
		if req.Password != req.ConfirmPassword {
			return errors.New("password and confirm password not match!")
		}
    }

    if err := uc.userRepo.UpdateUser(ctx, session.UserId, req); err != nil {
		if errors.Is(err, repository.ErrTodoAlreadyExists) {
			return repository.ErrTodoAlreadyExists
		}
		return fmt.Errorf("failed to add todo: %w", err)
	}
    
    var results string
    if len(update) > 1 {
        results = strings.Join(update, ", ")
    }

    action := fmt.Sprintf("User Update %s!", results)
	userLog := mapper.ToEntityFromNewLogUser(session.UserId, action)
	if err := uc.userLogRepo.UserLog(ctx, &userLog); err != nil {
		return fmt.Errorf("failed to save log: %w", err)
	}

	updateSession := mapper.ToEntityFromSession(session.UserId, session.Name, session.Username, session.Email, session.LoginTime)
	if err := uc.sessionRepo.Session(ctx, &updateSession); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}


    return nil
}