package ports

import "github.com/EvertonTomalok/ports-challenge/internal/core/domain"

type Repository interface {
	Upsert(key string, value domain.Port) error
	Get(key string) (domain.Port, bool)
	Size() int
}
