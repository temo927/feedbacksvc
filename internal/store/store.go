
package store

import (
    "context"

    "github.com/temo927/feedbacksvc/internal/domain"
)

type Store interface {
    Save(ctx context.Context, f domain.Feedback) error
}
