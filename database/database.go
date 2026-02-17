package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetTokenIssuedDate(id int) (time.Time, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(dbUser, dbPassword, dbName string) (*PostgresStorage, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Init() error {
	if err := s.createUserTable(); err != nil {
		return fmt.Errorf("failed to create user table: %v", err)
	}
	return nil
}

func (s *PostgresStorage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
