package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	_ "github.com/lib/pq"
)

var uniq_filter map[string]int

func isUnique(u User) bool {
	if _, ok := uniq_filter[u.Name]; !ok {
		uniq_filter[u.Name] = 1
		return true
	}
	return false
}

func EscapeQ(s string) string {
	return strings.ReplaceAll(s, "'", "\\'")
}

func FormQueryValues(u User) string {
	return fmt.Sprintf("('%s', '%s', '%s', '%s', %d, %d)",
		EscapeQ(u.Name),
		EscapeQ(u.Email),
		EscapeQ(u.Bio),
		EscapeQ(u.Birthdate),
		u.Gender,
		u.SexPref,
	)
}

// psql --user matcha --dbname matcha_db --host localhost
func InjectorMain(dsn, filename string) {
	// init unique map
	uniq_filter = make(map[string]int)

	// Read users json string
	data, err := ioutil.ReadFile(filename)
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
	// get struct from json string
	var users []User
	if json.Unmarshal(data, &users) != nil {
		panic(err)
	}
	// Connect to database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	// Forming single query
	query := `INSERT INTO users (user_nickname, user_email, bio, birthdate, gender, sexpref) VALUES`
	var sep string
	for i := 0; i < len(users); i++ {
		if isUnique(users[i]) {
			query += sep + FormQueryValues(users[i])
			sep = ","
		} else {
			fmt.Printf("Warning: non-unique user was skipped: Name='%s'\n", users[i].Name)
		}
	}
	query += ";"
	// query execution
	_, err = db.Exec(query)
	db.Close()
	if err != nil {
		panic(err)
	}
}
