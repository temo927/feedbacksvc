package stdout

import (
    "context"
    "encoding/json"
    "fmt"
    "time"

    "github.com/temo927/feedbacksvc/internal/pubsub"
)

type publisher struct{}

func New() pubsub.Publisher { return &publisher{} }

func (p *publisher) Publish(ctx context.Context, topic string, payload any) error {
    b, _ := json.Marshal(payload)
    // Simulate network delay and demonstrate non-blocking nature in logs
    fmt.Printf("[pubsub:%s] %s %s\n", topic, time.Now().UTC().Format(time.RFC3339), string(b))
    return nil
}
