package ports

import "github.com/EvertonTomalok/ports-challenge/internal/core/domain"

type Service interface {
	Upsert(ports domain.PortData) error
}
