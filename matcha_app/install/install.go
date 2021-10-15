package install

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

var _ResetSqlQuery string
var _CheckSqlQuery string

func init() {
	sql_replacer := getSqlParamReplacer()
	_ResetSqlQuery = sql_replacer.Replace(_SQL_DROP_SCRIPT + _SQL_SCHEMA_SCRIPT + _SQL_TYPES_SCRIPT + _SQL_TABLES_SCRIPT)
	_CheckSqlQuery = sql_replacer.Replace(_SQL_CHECKS_SCRIPT)
}

// Must be launched for installation of the application.
// If application is installed this does nothing.
// Any error which is happened here, is critical.
func InstallMatchaApplication(con *sql.DB, logger *log.Logger) (err error) {
	var db_set bool

	// check database status, if we can't establish connection return error few times
	if err = checkDBConnection(con); err == nil {
		logger.Printf("Install(): check query: \"%s\"", _CheckSqlQuery)
		db_set, err = isDBSet(con)
		// reset database tables if they doesn't exist
		if err == nil && !db_set {
			logger.Printf("Intstall(): database schema isn't found, resetting database with query: \"%s\"", _ResetSqlQuery)
			err = resetDBSchema(con)
		} else {
			logger.Printf("Install(): database schema exists, no need to reset")
		}
	}
	return err
}

func isDBSet(con *sql.DB) (bool, error) {
	result := false
	rows, err := con.Query(_CheckSqlQuery)
	if err == nil {
		result = rows.Next()
	}
	return result, err
}

func resetDBSchema(con *sql.DB) error {
	return executeSQLScript(con, _ResetSqlQuery)
}

// Try to ping database until it successful
// If it is still failling after few tries, return error
func checkDBConnection(con *sql.DB) error {
	var err error
	for ping_counter := 0; ; {
		if err = con.Ping(); err != nil {
			if ping_counter > 5 {
				return fmt.Errorf("db_ping err: %v", err)
			}
			ping_counter += 1
			time.Sleep(5 * time.Second)
		} else {
			return nil
		}
	}
}

// Execute multiline SQL Script
func executeSQLScript(con *sql.DB, full_query string) error {
	qrs := strings.Split(full_query, ";")
	for line, qr := range qrs {
		_, err := con.Exec(qr)
		if err != nil {
			return fmt.Errorf(
				"sql_script execution failed on line=%d, err=\"%v\", query=\"%s\"",
				line, err, qr,
			)
		}
	}
	return nil
}
