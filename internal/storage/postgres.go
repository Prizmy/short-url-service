package storage

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"short-url-service/pkg/shortener"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Post(url string) (string, error) {
	var shortURL string

	err := s.db.QueryRow("SELECT short_url FROM urls WHERE original_url = $1", url).Scan(&shortURL)
	if err == sql.ErrNoRows {
		for {
			shortURL = shortener.Generate()
			_, err = s.db.Exec("INSERT INTO urls (short_url, original_url) VALUES ($1, $2) ON CONFLICT (short_url) DO NOTHING", shortURL, url)
			if err == nil {
				break
			}
		}
	}
	return shortURL, err
}

func (s *PostgresStorage) Get(shortURL string) (string, error) {
	var originalURL string
	err := s.db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		return "", errors.New("URL not found")
	}
	return originalURL, nil
}
