package database

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/kasirku"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
