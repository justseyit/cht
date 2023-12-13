package pgsqlservice

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	Username string = "postgres"
	Password string = "root.seyit122"
	DBname   string = "chatservice"
	Host     string = "localhost"
	Port     string = "5432"
	SSLmode  string = "disable"
)

// PostgreSQL database instance
var pgdb *sql.DB

func ConnectToPostgreSQL() error {
	db, err := sql.Open("postgres", "postgres://"+Username+":"+Password+"@"+Host+":"+Port+"/"+DBname+"?sslmode="+SSLmode)
	pgdb = db
	return err
}
