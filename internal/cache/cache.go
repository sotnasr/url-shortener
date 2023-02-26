package cache

import (
	"context"
	"fmt"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error) // TODO: Verify if is necessary use generics.
	Set(ctx context.Context, key string, value string, ttl float64) error
}

type InMemory struct {
	cache map[string]string
}

// NewInMemoryCache is used to create a new instance of InMemory
func NewInMemoryCache() Cache {
	return &InMemory{
		cache: make(map[string]string),
	}
}

// Get Used to retrieve a cached value.
func (m *InMemory) Get(ctx context.Context, key string) (string, error) {
	value, ok := m.cache[key]
	if !ok {
		return "", fmt.Errorf("key '%s' was not found on cache, maybe the same was not available", key)
	}

	return value, nil
}

// Set Used to store values based on key on the cache context.
func (m *InMemory) Set(ctx context.Context, key string, value string, ttl float64) error {
	_, exists := m.cache[key]
	if exists {
		return fmt.Errorf("this key aready exits on cache")
	}

	m.cache[key] = value

	return nil
}
