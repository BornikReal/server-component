package storage_service

import (
	"fmt"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type storageRepository struct {
	storage Storage
}

func NewRepository(storage Storage) *storageRepository {
	return &storageRepository{
		storage: storage,
	}
}

func (r *storageRepository) get(key string) (string, error) {
	value, err := r.storage.Get(key)
	if err != nil {
		return "", fmt.Errorf("resposiotry.get: %w", err)
	}
	return value, nil
}

func (r *storageRepository) set(key string, value string) error {
	err := r.storage.Set(key, value)
	if err != nil {
		return fmt.Errorf("resposiotry.set: %w", err)
	}
	return nil
}
