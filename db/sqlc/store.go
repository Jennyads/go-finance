package db 

import "database/sql"

type Store interface {
	Querier 
	
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStor(db *sql.DB) Store{
	return &SQLStore{
		db: db,
		Queries: New(db),
	}
}