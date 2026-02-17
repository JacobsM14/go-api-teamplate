package db

import (
	"database/sql"
	"fmt"
	"time"
)

func (s *PostgresStorage) createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		display_name TEXT UNIQUE NOT NULL,
		last_token_issued TIMESTAMP,
		password_hash TEXT NOT NULL
	);`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) GetTokenIssuedDate(id int) (time.Time, error) {
	var lastTokenIssued time.Time

	err := s.db.QueryRow(
		`SELECT last_token_issued FROM users WHERE id = $1`,
		id,
	).Scan(&lastTokenIssued)

	if err != nil {
		if err == sql.ErrNoRows {
			// No user found with the given ID
			return time.Time{}, fmt.Errorf("unable to retrieve token information")
		}
		// Other database query errors
		return time.Time{}, fmt.Errorf("database error: %v", err)
	}

	return lastTokenIssued, nil
}
