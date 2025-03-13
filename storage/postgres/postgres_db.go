package storage_postgres

import (
	"fmt"
	"oopLab1/config"
	"sync"

	"github.com/jmoiron/sqlx"
)

type PostgresDB struct {
	DB *sqlx.DB
}

var once sync.Once
var dbInstance *PostgresDB

func GetPostgresDB(configuration *config.DBConfig) *PostgresDB {
	once.Do(func() {
		conn_string := fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s",
			configuration.User,
			configuration.Password,
			configuration.Host,
			configuration.Port,
			configuration.DBName,
		)

		db, err := sqlx.Open("pgx", conn_string)

		if err != nil {
			panic(fmt.Sprintf("Error connecting to database: %s", err.Error()))
		}

		dbInstance.DB = db
	})

	return dbInstance
}
