package repository

import (
    "errors"
    "context"

    "cli-todo-app-golang/internal/domain/entity"
)

var (
    ErrSessionNotFound = errors.New("session not found")
    ErrTodoNotFound = errors.New("todo not found or not yours")
)

type SessionRepository interface {
    Get(ctx context.Context) (*entity.Session, error)
    Session(ctx context.Context, newsession *entity.Session) error
    ClearSession(ctx context.Context) error
}