package repositories

import (
	"github.com/EvertonTomalok/ports-challenge/internal/domain"
)

func NewMemDB(options ...Option) *memDB {
	config := &config{maxLen: 1048576} // default max len is 2 ** 20
	for _, o := range options {
		o(config)
	}

	db := &memDB{data: make(map[string]domain.Port), maxSize: config.maxLen}
	return db
}

// In memory map struct database
type memDB struct {
	data    domain.PortData
	maxSize int
}

// Upsert the database ports collection.
func (repo *memDB) Upsert(key string, value domain.Port) error {
	// if value is not present on map but max size was achieve, it will
	// return error
	if len(repo.data) == repo.maxSize {
		_, found := repo.Get(key)
		if !found {
			return MaxSizeAchievedErr
		}
	}

	repo.data[key] = value
	return nil
}

func (repo *memDB) Get(key string) (domain.Port, bool) {
	item, found := repo.data[key]
	return item, found
}
