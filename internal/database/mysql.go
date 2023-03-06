package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewMysqlConnection(uri string) (*sql.DB, error) {
	// Connect to MySQL.
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return nil, err
	}

	// Set up important parts as was told by the documentation.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Return our database instance.
	return db, nil
}
