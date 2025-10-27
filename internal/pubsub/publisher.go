package pubsub

import "context"

type Publisher interface {
    Publish(ctx context.Context, topic string, payload any) error
}
