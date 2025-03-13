package storage

import (
	"github.com/jmoiron/sqlx"
)

type Database interface {
	GetDB() *sqlx.DB
}
