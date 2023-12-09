package postgres

import (
	"database/sql"
	"fmt"

	// postgres driver.
	_ "github.com/lib/pq"
	"github.com/otakakot/ninshow/internal/adapter/gateway"
)

var _ gateway.RDBFactory = (*Postgres)(nil)

type Postgres struct {
	*sql.DB
}

func New() *Postgres {
	return &Postgres{}
}

func (ps *Postgres) Of(dsn string) (*gateway.RDB, error) {
	client, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres: %w", err)
	}

	ps.DB = client

	if err := client.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	if _, err := client.Query("SELECT 1"); err != nil {
		return nil, fmt.Errorf("failed to query postgres: %w", err)
	}

	return &gateway.RDB{
		Client: client,
	}, nil
}
