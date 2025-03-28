package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/EvertonTomalok/ports-challenge/internal/core/domain"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type postgresDB struct {
	*sql.DB
}

func NewPostgresDB(options ...Option) *postgresDB {
	postgresURI := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Panic(err)
	}
	return &postgresDB{db}
}

// Upsert the database ports collection.
func (db *postgresDB) Upsert(key string, port domain.Port) error {
	query := `INSERT INTO ports (
		identifier, name, city, country, province, timezone, code, alias, regions, coordinates, unlocs
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	ON CONFLICT (identifier)
	DO UPDATE SET
		identifier = EXCLUDED.identifier,
		name = EXCLUDED.name,
		city = EXCLUDED.city,
		country = EXCLUDED.country,
		province = EXCLUDED.province,
		timezone = EXCLUDED.timezone,
		alias = EXCLUDED.alias,
		regions = EXCLUDED.regions,
		coordinates = EXCLUDED.coordinates,
		unlocs = EXCLUDED.unlocs;
	`

	_, err := db.Exec(query,
		key,
		port.Name,
		port.City,
		port.Country,
		port.Province,
		port.Timezone,
		port.Code,
		pq.Array(port.Alias),       // Convert alias slice to PostgreSQL array
		pq.Array(port.Regions),     // Convert regions slice to PostgreSQL array
		pq.Array(port.Coordinates), // Convert coordinates slice to PostgreSQL array
		pq.Array(port.Unlocs),      // Convert unlocs slice to PostgreSQL array
	)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (db *postgresDB) Get(key string) (domain.Port, bool) {
	var port domain.Port
	query := `
		SELECT
			identifier, name, city, country, province, timezone, code, alias, regions, coordinates, unlocs
		FROM
			ports
		WHERE
			identifier = $1;
	`

	// Execute the query and scan the result into the Port struct
	row := db.QueryRow(query, key)

	// If the result is found, fill the struct
	err := row.Scan(
		&port.Identifier,
		&port.Name,
		&port.City,
		&port.Country,
		&port.Province,
		&port.Timezone,
		&port.Code,
		pq.Array(&port.Alias),       // Scan the alias array
		pq.Array(&port.Regions),     // Scan the regions array
		pq.Array(&port.Coordinates), // Scan the coordinates array
		pq.Array(&port.Unlocs),      // Scan the unlocs array
	)

	if err != nil {
		fmt.Printf("%+v\n", err)
		if err == sql.ErrNoRows {
			return port, false
		}
		// todo: change the interface to return error instead bool
		return port, false
	}

	return port, true
}

func (db *postgresDB) Size() int {
	var count int
	query := `SELECT COUNT(*) FROM ports;`

	// Execute the query and scan the result into the count variable
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		fmt.Printf("error: %+v", err)
		return 0
	}

	return count
}
