package storage

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("value not found")
)

type Storage[K comparable, T any] struct {
	memory map[K]T

	mu sync.RWMutex
}

func NewStorage[K comparable, T any]() *Storage[K, T] {
	return &Storage[K, T]{memory: make(map[K]T)}
}

func (s *Storage[K, T]) Get(key K) (*T, error) {
	var r T

	s.mu.RLock()
	r, ok := s.memory[key]
	s.mu.RUnlock()

	if !ok {
		return nil, ErrNotFound
	}

	return &r, nil
}

func (s *Storage[K, T]) Set(key K, elem *T) {
	s.mu.Lock()
	s.memory[key] = *elem
	s.mu.Unlock()
}
