package storage

import (
	"github.com/jehaby/webapp102/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.C) (*sqlx.DB, error) {
	return sqlx.Open("postgres", cfg.DB.Conn)
}
