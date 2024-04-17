package driver

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(dsn string)(*DB, error){
	db, err := sql.Open("pgx", dsn)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	dbConn.SQL = db

	return dbConn, err
}