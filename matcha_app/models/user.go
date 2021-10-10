package models

import (
	"database/sql"
	"fmt"
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

	// This data is omited in json
	PassHash []byte `json:"-"`
	PassSalt []byte `json:"-"`
}

// TODO: maybe transfer everything to list of users
func (u *User) Insert() error {
	_, err := dbcon.Get().Exec(`
		INSERT INTO users
			(user_nickname, user_email, bio, birthdate, gender, sexpref, pass_hash, pass_salt)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
	`, u.Name, u.Email, u.Bio, u.BirthDate, u.Gender, u.SexPref, u.PassHash, u.PassSalt)

	return err
}

func (u *User) DeleteById() error {
	_, err := dbcon.Get().Exec(`DELETE FROM users WHERE user_id = $1`, u.Id)
	return err
}

type UserList []User

func (ulst_ptr *UserList) scanRows(rows *sql.Rows) error {
	var err error

	ulst := UserList(make([]User, 0))
	for rows.Next() && err == nil {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Bio, &u.BirthDate, &u.Gender, &u.SexPref, &u.PassHash, &u.PassSalt)
		ulst = append(ulst, u)
	}

	*ulst_ptr = ulst
	return err
}

func (ulst_ptr *UserList) selectWhereEq(db_key string, v interface{}) error {
	query := fmt.Sprintf(`
		SELECT user_id, user_nickname, user_email, bio, birthdate, gender, sexpref, pass_hash, pass_salt
		FROM users
		WHERE %s = $1
	`, db_key)

	rows, err := dbcon.Get().Query(query, v)

	if err == nil {
		err = ulst_ptr.scanRows(rows)
		rows.Close()
	}

	return err
}

func (ulst_ptr *UserList) SelectAll() error {

	rows, err := dbcon.Get().Query(`
	SELECT user_id, user_nickname, user_email, bio, birthdate, gender, sexpref, pass_hash, pass_salt
	FROM users
	`)

	if err == nil {
		err = ulst_ptr.scanRows(rows)
		rows.Close()
	}

	return err
}

//go:inline
func (us *UserList) SelectById(user_id int64) error {
	return us.selectWhereEq("user_id", &user_id)
}

//go:inline
func (us *UserList) SelectByName(user_name string) error {
	return us.selectWhereEq("user_nickname", &user_name)
}
