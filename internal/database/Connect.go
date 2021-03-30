package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// creates a SQL connection
// via go-sqlite3 connector
// and returns it
func Connect() *sql.DB {

	connStr := "root" + "@tcp(" + "127.0.0.1:3306" + ")/" + "analasia"
	conn, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err.Error())
	}
	return conn

}
