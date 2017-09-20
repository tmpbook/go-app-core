package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // dbDriver
)

// DBConn dataSourceName : "user:pwd@host:port/databasename"
func DBConn(dataSourceName string) (db *sql.DB) {

	dbDriver := "mysql" // Database driver

	// Realize the connection with mysql driver
	db, err := sql.Open(dbDriver, dataSourceName)

	// If error stop the application
	if err != nil {
		panic(err)
	}

	// Return db object to be used by another functions
	return db
}
