package store

import (
	"database/sql"
	"log"

	"github.com/goocarry/cnord-test/internal/config"
	//
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config          *config.Config
	log 			*log.Logger
	db              *sql.DB
	userRepository *UserRepository
}

// New ...
func New(config *config.Config, log *log.Logger) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	dsn := s.config.PostgresURL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() error {
	return s.db.Close()
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}