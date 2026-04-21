package repository

import (
    "context"
    "errors"
    
    "cli-todo-app-golang/internal/domain/entity"
)

var ErrUserLogNotFound = errors.New("user log not found!")

type UserLogRepository interface {
    UserLog(ctx context.Context, newlog *entity.UserLog) error
    GetAll(ctx context.Context) ([]entity.UserLog, error)
}