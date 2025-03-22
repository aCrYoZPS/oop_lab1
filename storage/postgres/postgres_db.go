package storage_postgres

import (
	"fmt"
	"oopLab1/config"
	"oopLab1/pkg/logger"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var once sync.Once
var dbInstance *sqlx.DB

func GetPostgresDB(configuration *config.DBConfig) *sqlx.DB {
	once.Do(func() {
		conn_string := fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
			configuration.User,
			configuration.Password,
			configuration.Host,
			configuration.Port,
			configuration.DBName,
			configuration.SSLMode,
		)

		db, err := sqlx.Open("postgres", conn_string)

		if err != nil {
			logger.Fatal(fmt.Sprintf("Error connecting to database: %s", err.Error()))
		}

		dbInstance = db
	})

	return dbInstance
}
