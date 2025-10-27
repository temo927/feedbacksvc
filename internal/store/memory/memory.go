package memory

import (
    "context"
    "sync"

    "github.com/temo927/feedbacksvc/internal/domain"
    "github.com/temo927/feedbacksvc/internal/store"
)

type memStore struct {
    mu   sync.RWMutex
    data map[string]domain.Feedback
}

func New() store.Store {
    return &memStore{
        data: make(map[string]domain.Feedback),
    }
}

func (m *memStore) Save(ctx context.Context, f domain.Feedback) error {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.data[f.ID] = f
    return nil
}
