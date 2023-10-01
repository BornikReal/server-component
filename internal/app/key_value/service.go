package key_value

import "fmt"

type repository interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type KeyValueService struct {
	rep repository
}

func NewKeyValueService(repository repository) *KeyValueService {
	return &KeyValueService{
		rep: repository,
	}
}

func (r *KeyValueService) Get(key string) (string, error) {
	value, err := r.rep.Get(key)
	if err != nil {
		return "", fmt.Errorf("KeyValueService.Get: %w", err)
	}
	return value, nil
}

func (r *KeyValueService) Set(key string, value string) error {
	err := r.rep.Set(key, value)
	if err != nil {
		return fmt.Errorf("KeyValueService.Set: %w", err)
	}
	return nil
}
