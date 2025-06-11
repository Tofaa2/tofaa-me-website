package sql

import (
	"database/sql"
)

type SqlDatabase interface {
	Db() *sql.DB
	Close() error
	Query(query string, args ...any) (*sql.Rows, error)
}
