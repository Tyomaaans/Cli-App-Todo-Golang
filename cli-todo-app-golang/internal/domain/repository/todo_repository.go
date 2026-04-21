package repository

import (
    "context"
    "errors"

    "cli-todo-app-golang/internal/domain/entity"
    "cli-todo-app-golang/internal/interface/dto"
)

var ErrTodoAlreadyExists = errors.New("todo already exists!")

type TodoRepository interface {
    GetAll(ctx context.Context) (map[string][]entity.Todo, error)

    AddTodo(ctx context.Context, todo *entity.Todo) error
    DeleteTodo(ctx context.Context, deleteId int, userId string) error
    GetTodoByUserId(ctx context.Context, userId string) ([]entity.Todo, error)
    UpdateTodo(ctx context.Context, userId string, req dto.UpdateTodoRequest) error
}