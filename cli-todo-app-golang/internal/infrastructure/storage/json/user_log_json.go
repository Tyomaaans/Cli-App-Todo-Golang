package json

import (
	"sync"
	"context"
	"encoding/json"
	"os"
	"errors"
	"fmt"

	"cli-todo-app-golang/internal/domain/entity"
	"cli-todo-app-golang/internal/domain/repository"
	"cli-todo-app-golang/internal/interface/dto"
	"cli-todo-app-golang/internal/interface/mapper"
)

type UserLogRepository struct {
    filePath string
	mu       sync.RWMutex
}

func NewLogRepository(filePath string) *UserLogRepository {
    return &UserLogRepository{
		filePath: filePath,
	}
}

func (r *UserLogRepository) load() ([]entity.UserLog, error) {
	fileData, err := os.ReadFile(r.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, repository.ErrUserLogNotFound
		}
		return nil, fmt.Errorf("failed to read user log: %w", err)
	}

	var storageLogUsers []dto.UserLogStorage
    if err := json.Unmarshal(fileData, &storageLogUsers); err != nil {
        return nil, fmt.Errorf("failed to parse user log: %w", err)
    }

	var logs []entity.UserLog
    for _, u := range storageLogUsers {
        logs = append(logs, mapper.ToEntityFromUserLogStorage(u))
    }

	return logs, nil
}

func (r *UserLogRepository) save(logs []entity.UserLog) error {
	var storageLogUsers []dto.UserLogStorage

	for _, u := range logs {
		storageLogUsers = append(storageLogUsers, mapper.ToUserLogStorageFromEntity(u))
	}

	data, err := json.MarshalIndent(storageLogUsers, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize user log: %w", err)
	}

	tmpFile := r.filePath + ".tmp"

	if err := os.WriteFile(tmpFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write temp user log file: %w", err)
	}

	if err := os.Rename(tmpFile, r.filePath); err != nil {
		_ = os.Remove(tmpFile)
		return fmt.Errorf("failed to finalize user log file: %w", err)
	}

	return nil
}

func (r *UserLogRepository) UserLog(ctx context.Context, newlog *entity.UserLog) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	logs, err := r.load()
	if err != nil {
		return err
	}

	logs = append(logs, *newlog)

	return r.save(logs)
}

func (r *UserLogRepository) GetAll(ctx context.Context) ([]entity.UserLog, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.load()
}