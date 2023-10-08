package inmemory

import (
	"fmt"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type StorageService struct {
	storage Storage
}

func NewStorageService(storage Storage) *StorageService {
	return &StorageService{
		storage: storage,
	}
}

func (r *StorageService) Get(key string) (string, error) {
	value, err := r.storage.Get(key)
	if err != nil {
		return "", fmt.Errorf("resposiotry.Get: %w", err)
	}
	return value, nil
}

func (r *StorageService) Set(key string, value string) error {
	err := r.storage.Set(key, value)
	if err != nil {
		return fmt.Errorf("resposiotry.Set: %w", err)
	}
	return nil
}
