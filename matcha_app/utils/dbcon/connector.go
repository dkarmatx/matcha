package dbcon

import (
	"database/sql"
	"matcha/config"

	_ "github.com/lib/pq"
)

// Database connector singleton
var c *sql.DB

// Get global DB connector singleton. Connection is created when this function called first time
func Get() *sql.DB {
	if c == nil {
		db, err := sql.Open(config.GetDBDriverName(), config.GetDSN())
		if err == nil {
			c = db
		} else {
			panic("Database Connector: Open() Error: " + err.Error())
		}
	}
	return c
}

// Close global DB connector
func Close() error {
	return c.Close()
}
