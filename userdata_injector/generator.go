package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Pallinder/go-randomdata"
)

func RandomGender() int {
	return GENDER_VAL(randomdata.Number(GENDER_MALE_IDX, GENDERS_AMOUNT))
}

func RandomGenderSet() int {
	return randomdata.Number(GENDER_MALE, GENDER_LIMIT_BIT)
}

func GenderToRandomDataGender(our_gender int) (their_gender int) {
	switch our_gender {
	case GENDER_MALE:
		their_gender = randomdata.Male
	case GENDER_FEMALE:
		their_gender = randomdata.Female
	case GENDER_OTHER:
		their_gender = randomdata.RandomGender
	default:
		panic("Incorrect gender value")
	}

	return
}

func GenerateUsers(n int) []User {
	var users []User
	for ; n > 0; n-- {
		gender := RandomGender()
		rgender := GenderToRandomDataGender(gender)
		u := User{
			Gender:    gender,
			SexPref:   RandomGenderSet(),
			Name:      randomdata.FirstName(rgender) + "_" + randomdata.LastName(),
			Birthdate: randomdata.FullDateInRange("1940-01-01", "2001-01-01"),
			Email:     randomdata.Email(),
			Bio:       randomdata.Paragraph(),
		}
		users = append(users, u)
	}
	return users
}

func GeneratorMain(n int, filename string) {
	users := GenerateUsers(n)
	if users_json_data, err := json.MarshalIndent(users, "", "    "); err != nil {
		panic(err)
	} else {
		ioutil.WriteFile(filename, users_json_data, 0644)
	}
}
