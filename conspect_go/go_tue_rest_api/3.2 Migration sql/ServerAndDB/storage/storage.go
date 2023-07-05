package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // для того что бы отработала функция init()
)

// Instance of storage
type Storage struct {
	config *Config
	// DataBase FileDescriptor
	db *sql.DB
}

// Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open connection
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfuly")
	return nil
}

// Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}
