package install

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"matcha/config"
	"strings"
	"text/template"
	"time"
)

func executeSQLScript(con *sql.DB, query_old string) error {
	qrs := strings.Split(query_old, ";")
	for line, qr := range qrs {
		_, err := con.Exec(qr)
		if err != nil {
			return errors.New(fmt.Sprintf(
				"SQL script failed on line=%d, error=\"%s\", query=\"%s\"",
				line,
				err.Error(),
				qr,
			))
		}
	}
	return nil
}

func proccessTemplate(tmpl_str string, vars interface{}) (string, error) {
	t, err := template.New("tmpl").Parse(tmpl_str)
	if err != nil {
		return "", err
	}

	var tmp bytes.Buffer
	err = t.Execute(&tmp, vars)
	if err != nil {
		return "", err
	}

	return tmp.String(), nil
}

var __SCHEMA_INIT_QUERY string

func init() {
	fmap := map[string]string{
		"Schema": config.GetDBSchemaName(),
	}
	// INSERT INTO users VALUES (2, 'Fedor', 'fedor@yandex.ru', '', '1999-02-19', 1, 2, decode('17023353311ba2c3067bf815e77b5ea72a6bd0c140907990ad2cce5087a8c175', 'hex'), 'abcd') ;
	// fedor123
	var err error
	__SCHEMA_INIT_QUERY, err = proccessTemplate(`
		DROP SCHEMA IF EXISTS {{.Schema}} CASCADE ;
		CREATE SCHEMA {{.Schema}} ;
		SET search_path TO {{.Schema}} ;

		CREATE DOMAIN {{.Schema}}.email_dom AS
		    varchar(1024)
		    CONSTRAINT valid_email CHECK (VALUE ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9_-]+([.][A-Za-z_-]+)+$');
		COMMENT ON DOMAIN email_dom IS 'Type of email with auto cheking';

		CREATE DOMAIN {{.Schema}}.gender_dom AS
		    int    DEFAULT 4
		    CONSTRAINT valid_gender CHECK(VALUE = 1 OR VALUE = 2 OR VALUE = 4);
		COMMENT ON DOMAIN gender_dom  IS '1 (001) - male, 2 (010) - female, 4 (100) - Other';

		CREATE DOMAIN {{.Schema}}.sexpref_dom AS
		    int    DEFAULT 7
		    CONSTRAINT valid_sexpref CHECK((VALUE & ~(7)) = 0);
		COMMENT ON DOMAIN sexpref_dom IS 'sexprefs is a bitset where each "gender-bit" is set if user is interested in matches with that gender';


		CREATE TABLE {{.Schema}}.users (
		    user_id         bigserial       UNIQUE NOT NULL,
		    user_nickname   varchar(255)    UNIQUE NOT NULL,
		    user_email      email_dom       UNIQUE NOT NULL,
		
		    bio             text            DEFAULT '',
		    birthdate       date            ,
		    gender          gender_dom      ,
		    sexpref         sexpref_dom     ,

		    pass_hash       bytea           ,
		    pass_salt       bytea           ,
		
		    PRIMARY KEY (user_id, user_nickname, user_email)
		);

		CREATE TABLE {{.Schema}}.auth_tokens (
		    value       bigint          UNIQUE NOT NULL,
		    expires_at  time            ,
		    user_id     bigint          ,

		    PRIMARY KEY (value),
		    FOREIGN KEY (user_id)       REFERENCES users (user_id) ON DELETE CASCADE
		);
		
		CREATE TABLE {{.Schema}}.user_connections (
		    conn_id     bigserial       UNIQUE NOT NULL,
		    user_id     bigint          NOT NULL,
		    friend_id   bigint          NOT NULL,
		
		    PRIMARY KEY (conn_id, user_id, friend_id),
		    FOREIGN KEY (user_id)       REFERENCES users (user_id) ON DELETE CASCADE,
		    FOREIGN KEY (friend_id)     REFERENCES users (user_id) ON DELETE CASCADE
		);
		
		CREATE TABLE {{.Schema}}.tags (
		    tag_id      bigserial       UNIQUE NOT NULL,
		    tag_name    varchar(255)    UNIQUE NOT NULL,
		
		    PRIMARY KEY (tag_id, tag_name)
		);
		
		CREATE TABLE {{.Schema}}.user_tags (
		    user_id     bigint          NOT NULL,
		    tag_id      bigint          NOT NULL,
		
		    PRIMARY KEY (user_id, tag_id),
		    FOREIGN KEY (user_id)       REFERENCES users (user_id) ON DELETE CASCADE,
		    FOREIGN KEY (tag_id)        REFERENCES tags (tag_id) ON DELETE CASCADE
		);
	`, fmap)

	if err != nil {
		panic("Install Matcha: init() error: " + err.Error())
	}
}

func isDBSet(con *sql.DB) (bool, error) {
	result := false
	rows, err := con.Query(`
		SELECT schema_name, schema_owner
			FROM information_schema.schemata
			WHERE
				schema_name = $1 AND schema_owner = $2 ;`,
		config.GetDBSchemaName(),
		config.GetDBConfig().User,
	)
	if err == nil {
		result = rows.Next()
	}
	return result, err
}

func resetDBSchema(con *sql.DB) error {
	return executeSQLScript(con, __SCHEMA_INIT_QUERY)
}

// Try to ping database until it successful
// If it is still failling after few tries, return error
func checkDBConnection(con *sql.DB) error {
	var err error
	for ping_counter := 0; ; {
		if err = con.Ping(); err != nil {
			if ping_counter > 5 {
				return errors.New("Could not establish DB connection: " + err.Error())
			}
			ping_counter += 1
			time.Sleep(5 * time.Second)
		} else {
			return nil
		}
	}
}

// Must be launched for installation of the application.
// If application is installed this does nothing.
// Any error which is happened here, is critical.
func InstallMatchaApplication(con *sql.DB, logger *log.Logger) error {
	var err error
	var db_set bool

	// check database status, if we can't establish connection return error few times
	if err = checkDBConnection(con); err != nil {
		return err
	}

	// reset database tables if they doesn't exist
	db_set, err = isDBSet(con)
	if err == nil && !db_set {
		logger.Printf("Database Schema is NOT found. Resetting database ...")
		err = resetDBSchema(con)
	}

	return err
}
