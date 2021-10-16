package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func EscapeQ(s string) string {
	return strings.ReplaceAll(s, "'", "\\'")
}

// psql --user matcha --dbname matcha_db --host localhost
func InjectorMain(dsn, filename string) {
	// Read users
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		panic(err)
	}
	// Connect to database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	// Insert values into tables
	query := `INSERT INTO users (user_nickname, user_email, bio, birthdate, gender, sexpref) VALUES`
	var sep string
	for i := 0; i < len(users); i++ {
		u := users[i]
		query += fmt.Sprintf("%s('%s', '%s', '%s', '%s', %d, %d)",
			sep,
			EscapeQ(u.Name),
			EscapeQ(u.Email),
			EscapeQ(u.Bio),
			EscapeQ(u.Birthdate),
			u.Gender,
			u.SexPref,
		)
		sep = ","
	}
	query += ";"
}
