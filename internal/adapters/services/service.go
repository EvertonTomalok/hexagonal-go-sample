package services

import (
	"github.com/EvertonTomalok/ports-challenge/internal/core/domain"
	"github.com/EvertonTomalok/ports-challenge/internal/ports"
)

type service struct {
	repository ports.Repository
}

func NewService(repository ports.Repository) ports.Service {
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
