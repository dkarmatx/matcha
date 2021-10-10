package models

import (
	"database/sql"
	"fmt"
	"matcha/auth"
	"matcha/dbcon"
	"time"
)

type AuthToken struct {
	Value     auth.UserToken `json:"token"`
	ExpiresAt time.Time      `json:"expires_at"`
	UserId    int64          `json:"user_id"`
}

func (tok *AuthToken) Insert() error {
	_, err := dbcon.Get().Exec(`
		INSERT INTO auth_tokens
			(value, expires_at, user_id)
		VALUES
			($1, $2, $3)
	`, tok.Value, tok.ExpiresAt, tok.UserId)

	return err
}

func (tok *AuthToken) UpdateByValue() error {
	_, err := dbcon.Get().Exec(`
		UPDATE auth_tokens
		SET value = $1, expires_at = $2, user_id = $3
		WHERE value = $1
	`, tok.Value, tok.ExpiresAt, tok.UserId)

	return err
}

func (tok *AuthToken) DeleteByValue() error {
	_, err := dbcon.Get().Exec(`DELETE FROM auth_tokens WHERE value = $1`, tok.Value)
	return err
}

func AuthTokenCreate(user_id int64, lifetime time.Duration) AuthToken {
	token := AuthToken{}
	token.UserId = user_id
	token.Value = auth.GenerateNewTokenValue()
	token.ExpiresAt = time.Now().Add(lifetime)
	return token
}

type AuthTokenList []AuthToken

func (toklst_ptr *AuthTokenList) scanRows(rows *sql.Rows) error {
	var err error

	toklst := AuthTokenList(make([]AuthToken, 0))
	for rows.Next() && err == nil {
		var t AuthToken
		err = rows.Scan(&t.Value, &t.ExpiresAt, &t.UserId)
		toklst = append(toklst, t)
	}

	*toklst_ptr = toklst
	return err
}

func (toklst_ptr *AuthTokenList) selectWhereEq(db_key string, v interface{}) error {
	query := fmt.Sprintf(`
		SELECT value, expires_at, user_id
		FROM auth_tokens
		WHERE %s = $1
	`, db_key)

	rows, err := dbcon.Get().Query(query, v)

	if err == nil {
		err = toklst_ptr.scanRows(rows)
		rows.Close()
	}

	return err
}

func (toklst_ptr *AuthTokenList) SelectByValue(value auth.UserToken) error {
	return toklst_ptr.selectWhereEq("value", value)
}

func (toklst_ptr *AuthTokenList) SelectByUserId(user_id int64) error {
	return toklst_ptr.selectWhereEq("user_id", user_id)
}
