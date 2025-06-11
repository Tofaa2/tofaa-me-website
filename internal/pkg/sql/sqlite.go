package sql

import "database/sql"

type SqliteDatabase struct {
	db *sql.DB
}

func (s *SqliteDatabase) Db() *sql.DB {
	return s.db
}

func (s *SqliteDatabase) Close() error {
	return s.db.Close()
}

func (s *SqliteDatabase) Query(query string, args ...any) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}

func NewSqlite(path string) (SqlDatabase, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &SqliteDatabase{db: db}, nil
}
