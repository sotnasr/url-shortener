package cache

import "context"

// Cache define the principal method s that can be used to interact
// with any implementation of cache interface.
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl float64) error
}
