package main

import (
	"os"
	"strconv"
)

type User struct {
	Name      string
	Birthdate string
	Bio       string
	Email     string
	Gender    int
	SexPref   int
}

const (
	GENDER_MALE_IDX = iota
	GENDER_FEMALE_IDX
	GENDER_OTHER_IDX
	GENDERS_AMOUNT

	GENDER_MALE   = 1 << GENDER_MALE_IDX
	GENDER_FEMALE = 1 << GENDER_FEMALE_IDX
	GENDER_OTHER  = 1 << GENDER_OTHER_IDX

	GENDER_LIMIT_BIT = 1 << GENDERS_AMOUNT
)

func GENDER_VAL(gender_idx int) int {
	return 1 << gender_idx
}

func main() {
	// GENERATOR: command  user_count       output_file
	// ./program  gen      10               users.json

	// INJECTOR:  command  db_dsn                                                                                  input_file
	// ./program  inject   "user=matcha dbname=matcha_db host=localhost port=5432 sslmode=disable password=m1234"  users.json
	if len(os.Args) < 4 {
		panic("invalid arguments")
	}

	command := os.Args[1]
	filename := os.Args[3]

	switch command {
	case "gen":
		if user_count, err := strconv.Atoi(os.Args[2]); err == nil {
			GeneratorMain(user_count, filename)
		} else {
			panic(err)
		}

	case "inject":
		dsn := os.Args[2]
		InjectorMain(dsn, filename)

	default:
		panic("invalid command")
	}

}
