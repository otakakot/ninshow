package gateway

import (
	"database/sql"
)

type RDB struct {
	Client *sql.DB
}

type RDBFactory interface {
	Of(dsn string) (*RDB, error)
}
