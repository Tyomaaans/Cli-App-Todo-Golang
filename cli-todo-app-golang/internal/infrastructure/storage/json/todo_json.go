package json

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "sync"
	"context"
    "time"

    "cli-todo-app-golang/internal/domain/entity"
	"cli-todo-app-golang/internal/domain/repository"
    "cli-todo-app-golang/internal/interface/dto"
	"cli-todo-app-golang/internal/interface/mapper"
)

type TodoRepository struct {
    filePath string
    mu       sync.RWMutex
}

func NewTodoRepository(filePath string) *TodoRepository {
    return &TodoRepository{filePath: filePath}
}

func (r *TodoRepository) load() (map[string][]entity.Todo, error) {
    fileData, err := os.ReadFile(r.filePath)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            return map[string][]entity.Todo{}, nil
        }
        return nil, fmt.Errorf("failed to read todo: %w", err)
    }

    var storageTodos []dto.TodoStorage
    if err := json.Unmarshal(fileData, &storageTodos); err != nil {
        return nil, fmt.Errorf("failed to parse todo: %w", err)
    }

    todos := make(map[string][]entity.Todo)
    for _, s := range storageTodos {
        userId, userTodos := mapper.ToEntityFromTodoStorage(s)
        todos[userId] = userTodos
    }

    return todos, nil
}

func (r *TodoRepository) save(todos map[string][]entity.Todo) error {
    var storageTodos []dto.TodoStorage
    for userId, userTodos := range todos {
        storageTodos = append(storageTodos, mapper.ToTodoStorageFromEntity(userId, userTodos))
    }

    data, err := json.MarshalIndent(storageTodos, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to serialize todo: %w", err)
    }

    tmpFile := r.filePath + ".tmp"
    if err := os.WriteFile(tmpFile, data, 0600); err != nil {
        return fmt.Errorf("failed to write temp todo file: %w", err)
    }

    if err := os.Rename(tmpFile, r.filePath); err != nil {
        _ = os.Remove(tmpFile)
        return fmt.Errorf("failed to finalize todo file: %w", err)
    }

    return nil
}

func (r *TodoRepository) GetAll(ctx context.Context) (map[string][]entity.Todo, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    return r.load()
}

func (r *TodoRepository) AddTodo(ctx context.Context, todo *entity.Todo) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    todos, err := r.load()
    if err != nil {
        return err
    }

    for _, t := range todos[todo.UserId] {
        if t.Task == todo.Task {
            return repository.ErrTodoAlreadyExists
        }
    }

    todos[todo.UserId] = append(todos[todo.UserId], *todo)

    return r.save(todos)
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, deleteIndex int, userId string) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    todos, err := r.load()
    if err != nil {
        return err
    }

    foundIndex := -1
    for i, todo := range todos[userId] {
        if i == (deleteIndex - 1) && todo.UserId == userId {
            foundIndex = i
            break
        }
    }

    if foundIndex == -1 {
        return repository.ErrTodoNotFound
    }

    todos[userId] = append(todos[userId][:foundIndex], todos[userId][foundIndex+1:]...)
    
    return r.save(todos)
}

func (r *TodoRepository) GetTodoByUserId(ctx context.Context, userId string) ([]entity.Todo, error) {
    todos, err := r.load()
    if err != nil {
        return nil, err
    }

    userTodos, ok := todos[userId]
    if !ok || len(userTodos) == 0 {
        return nil, repository.ErrTodoNotFound
    }

    return userTodos, nil
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, userId string, req dto.UpdateTodoRequest) error {
	r.mu.Lock()
    defer r.mu.Unlock()

    todos, err := r.load()
    if err != nil {
        return err
    }

	foundIndex := -1
    for i, todo := range todos[userId] {
        if i == (req.UpdateIndex - 1) && todo.UserId == userId {
            foundIndex = i
            break
        }
    }

    if foundIndex == -1 {
        return repository.ErrTodoNotFound
    }
    
    if req.Task != "" {
        todos[userId][foundIndex].Task = req.Task
    }
    if req.Priority != "" {
        todos[userId][foundIndex].Priority = req.Priority
    }
    if req.Progress != "" {
        todos[userId][foundIndex].Progress = req.Progress
    } 
    if req.DueDate != "" {
        dueDate, _ := time.Parse("2006-01-02", req.DueDate)
        todos[userId][foundIndex].DueDate = dueDate
    }

	return r.save(todos)
}