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

type SessionRepository struct {
	filePath string
	mu       sync.RWMutex
}

func NewSessionRepository(filePath string) *SessionRepository {
	return &SessionRepository{
		filePath: filePath,
	}
}

func (r *SessionRepository) load() (*entity.Session, error) {
	fileData, err := os.ReadFile(r.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, repository.ErrSessionNotFound
		}
		return nil, fmt.Errorf("failed to read session: %w", err)
	}

	var sessionStorage dto.SessionStorage
	if err := json.Unmarshal(fileData, &sessionStorage); err != nil {
		return nil, fmt.Errorf("failed to parse session: %w", err)
	}

	session := mapper.ToEntityFromSessionStorage(sessionStorage)
	
	return &session, nil
}

func (r *SessionRepository) save(session *entity.Session) error {
	sessionStorage := mapper.ToSessionStorageFromEntity(*session)

	data, err := json.MarshalIndent(sessionStorage, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize session: %w", err)
	}

	tmpFile := r.filePath + ".tmp"

	if err := os.WriteFile(tmpFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write temp session file: %w", err)
	}

	if err := os.Rename(tmpFile, r.filePath); err != nil {
		_ = os.Remove(tmpFile)
		return fmt.Errorf("failed to finalize session file: %w", err)
	}

	return nil
}

func (r *SessionRepository) Session(ctx context.Context, newSession *entity.Session) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.save(newSession)
}

func (r *SessionRepository) Get(ctx context.Context) (*entity.Session, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.load()
}

func (r *SessionRepository) ClearSession(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := os.Remove(r.filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("failed to clear session: %w", err)
	}

	return nil
}