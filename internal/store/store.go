
package store

import (
    "context"

    "github.com/yourname/feedbacksvc/internal/domain"
)

type Store interface {
    Save(ctx context.Context, f domain.Feedback) error
}
