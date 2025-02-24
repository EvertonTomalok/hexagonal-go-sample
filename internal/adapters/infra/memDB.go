package infra

import (
	"github.com/EvertonTomalok/ports-challenge/internal/core/domain"
)

func NewMemDB(options ...Option) *memDB {
	config := &config{maxLen: 1048576} // default max len is 2 ** 20
	for _, o := range options {
		o(config)
	}

	db := &memDB{data: make(map[string]domain.Port), maxSize: config.maxLen}
	return db
}

// In memory map struct database with a maxSize attribute to control
// max size of the database.
type memDB struct {
	data    domain.PortData
	maxSize int
}

// Upsert the database ports collection.
func (db *memDB) Upsert(key string, value domain.Port) error {
	// if value is not present on the map but max size was achieve, it will
	// return error
	if len(db.data) == db.maxSize {
		_, found := db.Get(key)
		if !found {
			return MaxSizeAchievedErr
		}
	}

	db.data[key] = value
	return nil
}

func (db *memDB) Get(key string) (domain.Port, bool) {
	item, found := db.data[key]
	return item, found
}

func (db *memDB) Size() int {
	return len(db.data)
}
