package gosession

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	Value      []uint8
	Expires_on time.Time
}

type SessionResponse struct {
	ID         string
	Value      string //TODO :should be a type of JSON
	Expires_on time.Time
}

func NewStoreFromOptions(s *SessionStoreOptions) *Store {
	db, err := sql.Open("mysql", s.DSN)
	if err != nil {
		log.Print(err.Error())
		log.Fatal("Could not open the database connection")
		return nil
	}
	return &Store{
		Db: db,
	}
}

func (s *Store) CreateSessionTable(tablename string) {
	tablename = "`" + strings.Trim(tablename, "`") + "`"
	query := "CREATE TABLE IF NOT EXISTS" + tablename + "(id VARCHAR(255) NOT NULL , session_data LONGBLOB , expires_on TIMESTAMP DEFAULT NOW(), PRIMARY KEY(`id`))"
	result, err := s.Db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	if rows, _ := result.RowsAffected(); rows > 0 {
		fmt.Println("Session table created")
	}
}

func (s *Store) CreateSession(tablename string, sessionId string, expires_on time.Time, session_data map[string]interface{}) {
	tablename = "`" + strings.Trim(tablename, "`") + "`"
	query := "INSERT INTO " + tablename + " (id , session_data , expires_on)  VALUES (?,?,?)"
	bs, _ := json.Marshal(session_data)
	_, err := s.Db.Exec(query, sessionId, bs, expires_on)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func (s *Store) GetSessionData(SessionId string) *SessionResponse {
	session := Session{}
	query := "SELECT * FROM sessions WHERE ID = ?"
	rows, err := s.Db.Query(query, SessionId)
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&session.ID, &session.Value, &session.Expires_on)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	val := string([]byte(session.Value))
	return &SessionResponse{
		ID:         session.ID,
		Value:      val,
		Expires_on: session.Expires_on,
	}

}

func (s *Store) DeleteExpiredSession(sessionId string) {
	data := s.GetSessionData(sessionId)
	query := "DELETE * FROM sessions WHERE ID = ? AND ? > NOW()"
	s.Db.Query(query, sessionId, data.Expires_on)
}

func (s *Store) StartCleanUP(SessionId string) {
	go RunCleanUp(s, SessionId)
}
