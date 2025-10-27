package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/temo927/feedbacksvc/internal/domain"
	"github.com/temo927/feedbacksvc/internal/store"
)

type fsStore struct {
	client     *firestore.Client
	collection string
}

func New(ctx context.Context, projectID, collection string) (store.Store, error) {
	if collection == "" {
		collection = "feedback"
	}
	c, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &fsStore{client: c, collection: collection}, nil
}

func (s *fsStore) Save(ctx context.Context, f domain.Feedback) error {
	// Write with a small timeout so we don’t hang the handler
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := s.client.Collection(s.collection).Doc(f.ID).Set(cctx, map[string]any{
		"id":         f.ID,
		"name":       f.Name,
		"email":      f.Email,
		"message":    f.Message,
		"created_at": f.CreatedAt,
	})
	return err
}
