package ports

import "github.com/EvertonTomalok/ports-challenge/internal/core/domain"

//go:generate mockgen -source=./service.go -destination=./service_mock.go -package=ports
type Service interface {
	Upsert(ports domain.PortData) error
}
