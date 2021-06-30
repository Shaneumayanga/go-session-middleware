package gosession

import (
	"database/sql"
	"log"
)

type SessionStoreOptions struct {
	DSN       string
	TableName string
}

type Store struct {
	Db *sql.DB
}

func NewStoreFromOptions(s *SessionStoreOptions) *Store {
	db, err := sql.Open("mysql", s.DSN)
	if err != nil {
		log.Fatal("Could not open the database connection")
	}
	return &Store{
		Db: db,
	}
}

func (s *Store) CreateSessionTable(tablename string) {
	query := "CREATE TABLE IF NOT EXISTS" + tablename + "(id INT NOT NULL AUTO_INCREMENT , session_data LONGBLOB , expires_on TIMESTAMP DEFAULT NOW(), PRIMARY KEY(`id`))"
	_, err := s.Db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Store) CreateSession(sessionId string, tablename string) {

}
