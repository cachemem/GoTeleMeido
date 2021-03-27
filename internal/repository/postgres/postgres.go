package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type pgRepo struct {
	db *sql.DB
}

func New(dbUrl string) *pgRepo {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or another initialization error.
		log.Fatal(err)
	}
	return &pgRepo{db: db}
}

func (pgr *pgRepo) GetOwner() int64 {
	var ownerId int64
	rows, err := pgr.db.Query("SELECT id FROM users WHERE is_owner = $1 LIMIT 1", true)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ownerId)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ownerId
}

func (pgr *pgRepo) Close() {
	pgr.db.Close()
}
