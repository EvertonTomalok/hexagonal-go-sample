package ports

import "github.com/EvertonTomalok/ports-challenge/internal/domain"

type Service interface {
	Upsert(ports domain.PortData) error
	Get(key string) (domain.Port, bool)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

// Upsert will insert or update each key/value contained in the map received in the function.
func (s *service) Upsert(ports domain.PortData) error {
	for key, value := range ports {
		if err := s.repository.Upsert(key, value); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) Get(key string) (domain.Port, bool) {
	return s.repository.Get(key)
}
