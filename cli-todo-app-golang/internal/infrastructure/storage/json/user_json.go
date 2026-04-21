package json

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"cli-todo-app-golang/internal/domain/entity"
	"cli-todo-app-golang/internal/domain/repository"
	"cli-todo-app-golang/internal/interface/dto"
	"cli-todo-app-golang/internal/interface/mapper"
)

type UserRepository struct {
	filePath string
	mu       sync.RWMutex
}

func NewUserRepository(filePath string) *UserRepository {
	return &UserRepository{filePath: filePath}
}

func (r *UserRepository) load() ([]entity.User, error) {
	fileData, err := os.ReadFile(r.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []entity.User{}, nil
		}
		return nil, fmt.Errorf("failed to read user: %w", err)
	}

	var storageUsers []dto.UserStorage
    if err := json.Unmarshal(fileData, &storageUsers); err != nil {
        return nil, fmt.Errorf("failed to parse user: %w", err)
    }

    var users []entity.User
    for _, u := range storageUsers {
        users = append(users, mapper.ToEntityFromUserStorage(u))
    }

	return users, nil
}

func (r *UserRepository) save(users []entity.User) error {
	var storageUsers []dto.UserStorage

	for _, u := range users {
		storageUsers = append(storageUsers, mapper.ToUserStorageFromEntity(u))
	}

	data, err := json.MarshalIndent(storageUsers, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize user: %w", err)
	}

	tmpFile := r.filePath + ".tmp"

	if err := os.WriteFile(tmpFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write temp user file: %w", err)
	}

	if err := os.Rename(tmpFile, r.filePath); err != nil {
		_ = os.Remove(tmpFile)
		return fmt.Errorf("failed to finalize user file: %w", err)
	}

	return nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.load()
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users, err := r.load()
	if err != nil {
		return nil, err
	}

	for i := range users {
		if users[i].Username == username {
			return &users[i], nil
		}
	}

	return nil, repository.ErrUserNotFound
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users, err := r.load()
	if err != nil {
		return nil, err
	}

	for i := range users {
		if users[i].Email == email {
			return &users[i], nil
		}
	}

	return nil, repository.ErrUserNotFound
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	users, err := r.load()
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Username == user.Username || u.Email == user.Email {
			return repository.ErrUserAlreadyExists
		}
	}

	users = append(users, *user)

	return r.save(users)
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	users, err := r.load()
	if err != nil {
		return err
	}

	for i := range users {
		if users[i].UserId == user.UserId {
			users[i] = *user
			return r.save(users)
		}
	}

	return repository.ErrUserNotFound
}

func (r *UserRepository) ForgotPassword(ctx context.Context, email, password string) error {
	user, err := r.GetByEmail(ctx, email)
	if err != nil {
        return err
    }

	user.Password = password

	return r.Update(ctx, user)
}

func (r *UserRepository) UpdateName(ctx context.Context, username, name string) error {
	user, err := r.GetByUsername(ctx, username)
	if err != nil {
        return err
    }

	user.Name = name

	return r.Update(ctx, user)
}

func (r *UserRepository) UpdateEmail(ctx context.Context, username, email string) error {
	user, err := r.GetByUsername(ctx, username)
	if err != nil {
        return err
    }

	user.Email = email

	return r.Update(ctx, user)
}

func (r *UserRepository) UpdateUsername(ctx context.Context, username, newusername string) error {
	user, err := r.GetByUsername(ctx, username)
	if err != nil {
        return err
    }

	user.Username = newusername

	return r.Update(ctx, user)
}

func (r *UserRepository) UpdateUser(ctx context.Context, userId string, req dto.UpdateUserRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	users, err := r.load()
	if err != nil {
		return err
	}

	foundIndex := -1
	for i := range users {
		if users[i].UserId == userId {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return repository.ErrUserNotFound
	}

	if req.Name != "" {
		users[foundIndex].Name = req.Name
	}

	if req.Username != "" {
		users[foundIndex].Username = req.Username
	}

	if req.Email != "" {
		users[foundIndex].Email = req.Email
	}

	if req.Password != "" {
		users[foundIndex].Password = req.Password
	}

	return r.save(users)
}

func (r *UserRepository) DeleteUser(ctx context.Context, userId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	users, err := r.load()
	if err != nil {
		return err
	}

	foundIndex := -1
	for i := range users {
		if users[i].UserId == userId {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return repository.ErrUserNotFound
	}

	users = append(users[:foundIndex], users[foundIndex+1:]...)

	return r.save(users)

}