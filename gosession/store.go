package gosession

import "database/sql"

type SessionStoreOptions struct {
	Db        string
	TableName string
}

type Store struct {
	Db *sql.DB
}

func NewStoreFromOptions(s *SessionStoreOptions) *Store {
	//connection := sql.Open(s.DB)
	//Connection.exec("CREATE TABLE {tablename} (ID )")
	return &Store{}
}

func (s *Store) CreateSession(sessionId string, tablename string) {

}
