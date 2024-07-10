package storage

import (
	"auth_service/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectionDb() (*sql.DB, error) {
	cnf := config.Load()
	conDb := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", cnf.PostgresHost, cnf.PostgresPort, cnf.PostgresUser, cnf.PostgresDatabase, cnf.PostgresPassword)
	return sql.Open("postgres", conDb)
}
