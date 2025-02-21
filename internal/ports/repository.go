package ports

import "github.com/EvertonTomalok/ports-challenge/internal/domain"

type Repository interface {
	Upsert(key string, value domain.Port) error
	Get(key string) domain.Port
}
