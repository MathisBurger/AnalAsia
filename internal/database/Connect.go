package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// creates a SQL connection
// via go-sqlite3 connector
// and returns it
func Connect() *sql.DB {

	var connStr string

	if os.Getenv("mode") == "dev" {
		connStr = "root" + "@tcp(" + "127.0.0.1:3306" + ")/" + "analasia"
	} else {
		connStr = os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ")/" + os.Getenv("DATABASE_NAME")
	}
	conn, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err.Error())
	}
	return conn

}
