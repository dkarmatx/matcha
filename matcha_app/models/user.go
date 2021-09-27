package models

import (
	"matcha/dbcon"
)

/*
 *	SelectByColumnName() - load by some column (WHERE)
 *  SelectAll() - load all instances (if list model)
 *  Insert() - insert new instance (default, usualy without serial values or auto computed values)
 *	InsertWithColumnName() - insert by specified column value (for example id)
 *  Update() - Upate value
 *  DeleteByColumnName() - don't
 */

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	BirthDate string `json:"birth-date"`
	Gender    int    `json:"gender"`
	SexPref   int    `json:"sex_preferences"`
}

func (u *User) Insert() error {
	_, err := dbcon.Get().Exec(`
		INSERT INTO users
			(user_nickname, user_email, bio, birthdate, gender, sexpref)
		VALUES
			($1, $2, $3, $4, $5, $6)
	`, u.Name, u.Email, u.Bio, u.BirthDate, u.Gender, u.SexPref)

	return err
}

func (u *User) DeleteById() error {
	_, err := dbcon.Get().Exec(`DELETE FROM users WHERE user_id = $1`, u.Id)
	return err
}

type UserList []User

func (ulst_ptr *UserList) SelectAll() error {
	ulst := UserList(make([]User, 0))

	rows, err := dbcon.Get().Query(`
		SELECT user_id, user_nickname, user_email, bio, birthdate, gender, sexpref
		FROM users
	`)

	for rows.Next() && err == nil {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Bio, &u.BirthDate, &u.Gender, &u.SexPref)
		ulst = append(ulst, u)
	}
	rows.Close()

	*ulst_ptr = ulst
	return err
}
