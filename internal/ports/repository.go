package ports

import "github.com/EvertonTomalok/ports-challenge/internal/core/domain"

//go:generate mockgen -source=./repository.go -destination=./repository_mock.go -package=ports
type Repository interface {
	Upsert(key string, value domain.Port) error
	Get(key string) (domain.Port, bool)
	Size() int
}
