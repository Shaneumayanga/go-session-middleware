package gosession

import (
	"database/sql"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SessionStoreOptions struct {
	DSN       string
	TableName string
}

type Store struct {
	Db *sql.DB
}

type Session struct {
	ID         string
	Value      map[string]interface{}
	Expires_on time.Time
}

func NewStoreFromOptions(s *SessionStoreOptions) *Store {
	db, err := sql.Open("mysql", s.DSN)
	if err != nil {
		log.Print(err.Error())
		log.Fatal("Could not open the database connection")
	}
	return &Store{
		Db: db,
	}
}

func (s *Store) CreateSessionTable(tablename string) {
	tablename = "`" + strings.Trim(tablename, "`") + "`"
	query := "CREATE TABLE IF NOT EXISTS" + tablename + "(id VARCHAR(255) NOT NULL , session_data LONGBLOB , expires_on TIMESTAMP DEFAULT NOW(), PRIMARY KEY(`id`))"
	_, err := s.Db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Store) CreateSession(sessionId string, tablename string) {
	//query := "INSERT INTO"
}
